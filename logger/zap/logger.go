package zap

import (
	"github.com/aomi-go/common/logger"
	"go.uber.org/zap"
)

func NewZapLogger(provider *zap.Logger) *Logger {
	l := &Logger{
		BaseLogger: &logger.BaseLogger{
			NewLog: func() logger.Log {
				return Log{
					provider: provider,
					fields:   make([]zap.Field, 0),
				}
			},
		},
	}
	return l
}

type Logger struct {
	*logger.BaseLogger
}
