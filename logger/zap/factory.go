package zap

import (
	"github.com/aomi-go/common/logger"
	"go.uber.org/zap"
)

func NewZapLoggerCreator(cfg zap.Config, options ...zap.Option) (logger.Creator, error) {
	newOptions := []zap.Option{
		zap.AddCallerSkip(2),
	}

	newOptions = append(newOptions, options...)

	provider, err := cfg.Build(newOptions...)
	if nil != err {
		return nil, err
	}

	return func(name string) logger.Logger {
		l := provider.Named(name)
		return NewZapLogger(l)
	}, nil
}
