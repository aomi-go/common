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

2. 注册