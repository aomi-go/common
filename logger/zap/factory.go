package zap

import (
	"github.com/aomi-go/common/logger"
	"go.uber.org/zap"
)

func NewZapLoggerFactory(cfg zap.Config, options ...zap.Option) (*LoggerFactory, error) {
	newOptions := []zap.Option{
		zap.AddCallerSkip(2),
	}

	newOptions = append(newOptions, options...)

	provider, err := cfg.Build(newOptions...)
	if nil != err {
		return nil, err
	}
	return &LoggerFactory{
		provider: provider,
	}, nil
}

type LoggerFactory struct {
	provider *zap.Logger
}

func (z *LoggerFactory) Get(name string) logger.Logger {
	provider := z.provider.Named(name)
	return NewZapLogger(provider)
}

func (z *LoggerFactory) GetProvider() interface{} {
	return z.provider
}
