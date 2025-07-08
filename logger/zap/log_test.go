package zap

import (
	"fmt"
	"github.com/aomi-go/common/logger"
	"go.uber.org/zap"
	"testing"
)

func TestLog(t *testing.T) {

	factory, err := NewZapLoggerFactory(zap.NewDevelopmentConfig(), zap.Fields(zap.String("env", "dev")))
	if err != nil {
		t.Fatal(err)
	}
	logger.SetFactory(factory)
	log := logger.GetLogger("test")
	fmt.Printf("log的地址: %p\n", &log) // 输出: 0xc00001a0a8

	log.String("name", "Sean").Info("hello")
	prodFactory, err := NewZapLoggerFactory(zap.NewProductionConfig(), zap.Fields(zap.String("env", "prod")))
	if err != nil {
		t.Fatal(err)
	}
	logger.SetFactory(prodFactory)
	logger.ResetLogger()
	log.String("name", "Sean").Info("hello")
}
