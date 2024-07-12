# Logger

golang 日志API库；封装标准API，由其他日志库实现底层逻辑，使得使用方式得到统一、方便日期框架切换

### 目前实现的日志库

- [x] [zap](https://github.com/uber-go/zap)
- [ ] [logrus]()

### 使用方式

1. 安装
    ```shell
    go get -u github.com/aomi-go/common/logger
    # 根据自己的需求安装底层实现库, 例如，使用zap
    go get -u go.uber.org/zap
    ```

2. 使用
   ```go
   // logger/log.go 文件 初始化日志
   package logger

   import (
       "github.com/aomi-go/common/logger"
       "go.uber.org/zap"
   )

   // Init 初始化日志
   func Init() {
        cfg := zap.NewProductionConfig() 
        //cfg := zap.NewDevelopmentConfig()
        // 自定义配置 
        //cfg := zap.Config{}
   
        // 自定义日志创建
        //logger.Factory.SetCreator(func(name string) logger.Logger { 
        //    //provider, _ := zap.NewProduction()
        //    //provider, _ := zap.NewDevelopment()
        //    provider, _ := cfg.Build()
        //    provider = provider.Named(name)
        //    return logger.NewZapLogger(provider)
        //})
   
        // 使用默认创建起
        logger.Factory.SetCreator(logger.CreateZapCreator(cfg))
   }

   ```
   ```go
   // main.go 文件
   
   func main(){
        logger.Init()
        log := logger.Factory.GetLogger("test")
		
        log.Info().String("addr", "0.0.0.0:80").Msg("Start web server")
   }

   ```