package logger

import (
	"sync"
)

type Factory interface {
	Get(name string) Logger

	GetProvider() interface{}
}

var (
	factory Factory
	loggers = sync.Map{}
)

func SetFactory(f Factory) {
	factory = f
}

func GetLogger(name string) Logger {
	// 直接尝试从 sync.Map 中加载
	if logger, ok := loggers.Load(name); ok {
		return logger.(Logger)
	}

	if nil == factory {
		panic("logger factory is nil, please set it first: logger.SetCreator()")
	}

	newLogger := factory.Get(name)
	// 使用 LoadOrStore 避免并发创建相同 key 的 Logger
	if existing, loaded := loggers.LoadOrStore(name, newLogger); loaded {
		return existing.(Logger) // 已有其他 goroutine 抢先创建
	}

	return newLogger // 本次创建成功
}

func ResetLogger() {
	if nil == factory {
		panic("logger factory is nil, please set it first: logger.SetFactory()")
	}

	// 遍历 sync.Map 并重置每个 Logger
	loggers.Range(func(key, value interface{}) bool {
		logger := value.(Logger)
		logger.SetProvider(factory.GetProvider())
		return true // 继续遍历
	})
}
