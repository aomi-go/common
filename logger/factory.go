package logger

import (
	"sync"
)

type Creator func(name string) Logger

func defaultLoggerCreator(name string) Logger {
	return EmptyLogger
}

var (
	loggers = sync.Map{}
	creator = defaultLoggerCreator
)

func SetCreator(c Creator) {
	creator = c
	ResetLogger()
}

func GetLogger(name string) Logger {
	// 直接尝试从 sync.Map 中加载
	if logger, ok := loggers.Load(name); ok {
		return logger.(Logger)
	}
	newLogger := creator(name)
	loggerDelegate := NewLoggerDelegate(newLogger)
	// 使用 LoadOrStore 避免并发创建相同 key 的 Logger
	if existing, loaded := loggers.LoadOrStore(name, loggerDelegate); loaded {
		return existing.(Logger) // 已有其他 goroutine 抢先创建
	}

	return loggerDelegate // 本次创建成功
}

func ResetLogger() {
	// 遍历 sync.Map 并重置每个 Logger
	loggers.Range(func(key, value interface{}) bool {
		delegate := value.(*Delegate)
		delegate.Logger = creator(key.(string))
		return true // 继续遍历
	})
}
