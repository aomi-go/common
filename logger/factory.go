package logger

type Factory interface {
	Get(name string) Logger
	// Reset 重置已知logger
	Reset(loggers map[string]Logger)
}

var (
	factory Factory
	loggers = make(map[string]Logger)
)

func SetFactory(f Factory) {
	factory = f
}

func GetLogger(name string) Logger {
	if logger, ok := loggers[name]; ok {
		return logger
	}

	if nil == factory {
		panic("logger factory is nil, please set it first: logger.SetCreator()")
	}

	l := factory.Get(name)
	loggers[name] = l
	return l
}
func ResetLogger() {
	if nil == factory {
		panic("logger factory is nil, please set it first: logger.SetFactory()")
	}
	factory.Reset(loggers)
}
