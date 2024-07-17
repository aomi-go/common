package zap

import (
	"github.com/aomi-go/common/logger"
	"go.uber.org/zap"
	"testing"
)

func TestLog(t *testing.T) {

	logger.SetFactory(NewZapLoggerFactory(zap.NewDevelopmentConfig()))
	log := logger.GetLogger("test")
	log.Info("dev")
	logger.SetFactory(NewZapLoggerFactory(zap.NewProductionConfig()))
	logger.ResetLogger()
	log.Info("prod")
}
