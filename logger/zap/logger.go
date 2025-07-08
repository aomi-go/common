package zap

import (
	"github.com/aomi-go/common/logger"
	"go.uber.org/zap"
)

func NewZapLogger(provider *zap.Logger) *Logger {
	l := &Logger{
		BaseLogger: &logger.BaseLogger{},
	}
	l.SetProvider(provider)
	return l
}

type Logger struct {
	*logger.BaseLogger
	provider *zap.Logger
}

func (l *Logger) SetProvider(provider interface{}) {
	l.provider = provider.(*zap.Logger)
	l.NewLog = func() logger.Log {
		return Log{
			fields:   make([]zap.Field, 0),
			provider: l.provider,
		}
	}
}
