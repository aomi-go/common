package zap

import (
	"github.com/aomi-go/common/logger"
	"go.uber.org/zap"
)

func NewZapLoggerFactory(cfg zap.Config) *LoggerFactory {
	return &LoggerFactory{
		cfg: cfg,
	}
}

type LoggerFactory struct {
	cfg zap.Config
}

func (z *LoggerFactory) SetCfg(cfg zap.Config) {
	z.cfg = cfg
}

func (z *LoggerFactory) Get(name string) logger.Logger {
	provider := z.get(name)
	return NewZapLogger(provider)
}

func (z *LoggerFactory) Reset(loggers map[string]logger.Logger) {
	for name, l := range loggers {
		if zl, ok := l.(*Logger); ok {
			zl.provider = z.get(name)
		}
	}
}

func (z *LoggerFactory) get(name string) *zap.Logger {
	provider, err := z.cfg.Build()
	if nil != err {
		panic(err)
	}
	provider = provider.Named(name)
	defer provider.Sync()
	return provider
}
