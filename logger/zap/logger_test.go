package zap

import (
	"github.com/aomi-go/common/logger"
	"go.uber.org/zap"
	"testing"
)

func TestZapLog(t *testing.T) {

	log := logger.GetLogger("test")
	log.String("name", "Sean").Info("empty log provider")
	creator, err := NewZapLoggerCreator(zap.NewDevelopmentConfig(), zap.Fields(zap.String("env", "dev")))
	if err != nil {
		t.Fatal(err)
	}
	logger.SetCreator(creator)
	log.String("name", "Sean").Info("hello")

	creator, err = NewZapLoggerCreator(zap.NewProductionConfig(), zap.Fields(zap.String("env", "prod")))
	if err != nil {
		t.Fatal(err)
	}
	logger.SetCreator(creator)
	log.String("name", "Sean").Info("hello")
}
