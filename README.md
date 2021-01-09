```
 ██████╗  ██████╗        ██████╗ ██╗███╗   ██╗       █████╗ ██████╗ ██╗
██╔════╝ ██╔═══██╗      ██╔════╝ ██║████╗  ██║      ██╔══██╗██╔══██╗██║
██║  ███╗██║   ██║█████╗██║  ███╗██║██╔██╗ ██║█████╗███████║██████╔╝██║
██║   ██║██║   ██║╚════╝██║   ██║██║██║╚██╗██║╚════╝██╔══██║██╔═══╝ ██║
╚██████╔╝╚██████╔╝      ╚██████╔╝██║██║ ╚████║      ██║  ██║██║     ██║
 ╚═════╝  ╚═════╝        ╚═════╝ ╚═╝╚═╝  ╚═══╝      ╚═╝  ╚═╝╚═╝     ╚═╝

* [register -env fat]
* [register cors]
* [register rate]
* [register panic notify]
* [register pprof]
* [register swagger]
* [register prometheus]
```

## go-gin-api

基于 [Gin](https://github.com/gin-gonic/gin) 进行模块化设计的 API 框架，封装了常用的功能，使用简单，致力于进行快速的业务研发，同时增加了更多限制，以约束项目组开发成员、规避混乱无序和自由随意。

亮点功能：

- :star: [trace] 开发调试的辅助工具。

可记录如下信息：

- trace_id，当前请求的唯一ID
- request，当前请求的请求信息
- response，当前请求的返回信息
- third_party_requests，当前请求涉及到调用第三方的信息
- debugs，当前请求的调试信息
- sqls，当前请求执行的 sql 信息
- success，当前请求结果
- cost_seconds，执行时长，单位秒

供参考学习，线上使用请谨慎！

**查看 Jaeger 链路追踪代码，请查看 [v1.0版](https://github.com/xinliangnote/go-gin-api/releases/tag/v1.0)，文档 [jaeger.md](https://github.com/xinliangnote/go-gin-api/blob/master/docs/jaeger.md) 点这里**。

持续更新... 

## Catalogue

```cassandraql
├── cmd                          # 项目入口文件，api/main.go 为启动 HTTP API 服务
├── configs                      # 配置文件统一存放目录
├── docs                         # Swagger 文档，执行 swag init 生成的
├── init                         # 项目初始脚本，比如初始化 SQL 语句
├── internal                     # 业务目录
│   ├── api                      # 业务代码
│   ├── core                     # 脚本代码
│   ├── pkg                      # 内部使用的 package
├── logs                         # 存放日志的目录
└── pkg                          # 一些封装好的 package
```
## Features

- [x] 包管理工具 [Go Modules](https://github.com/golang/go/wiki/Modules)
- [x] [Gin](https://github.com/gin-gonic/gin) 支持优雅关闭
- [x] 配置文件解析库 [Viper](https://github.com/spf13/viper)
- [x] 文档使用 [Swagger](https://swagger.io/) 生成
- [x] 性能分析使用 [pprof](github.com/gin-contrib/pprof)
- [x] [zap](go.uber.org/zap) 日志记录
- [x] [rate](golang.org/x/time/rate) 限流
- [x] [token] 基于[JWT](github.com/dgrijalva/jwt-go) 身份认证
- [x] [notify] 异常捕获并进行邮件告警
- [x] [trace] 开发调试的辅助工具
- [x] [errno] 统一定义错误码
- [x] [env] 支持 FAT、UAT、PRO 环境
- [x] [aes] AES 对称加密
- [x] [rsa] RSA 非对称加密
- [x] 数据库组件使用 [GORM V2](gorm.io/gorm)
- [x] Redis 组件使用 [go-redis](https://github.com/go-redis/redis)
- [ ] MongoDB
- [ ] Prometheus
- [ ] 任务调度
- [ ] gRPC
- [ ] ...

## Quick start

#### Requirements

- Go version >= 1.15
- Global environment configure (Linux/Mac)

```
export GO111MODULE=on
export GOPROXY=https://goproxy.io
```

#### Environment

```
-env fat

// dev 开发环境
// fat 测试环境[默认]
// uat 预发布环境
// pro 正式环境
```

#### Configs

配置文件目录：`./configs`，根据不同的环境变量使用不同的配置文件。

项目启动时，需配置如下配置：

- MySQL 配置，主、从和基础项；
- Redis 配置；
- Mail 配置；

#### Init

项目初始化目录：`./init`。

- db/tables.sql，初始化 MySQL 表结构；

#### Build & Run

```
cd go-gin-api

go run main.go

// 启动成功后，可访问：http://127.0.0.1:9999/h/info
```

#### swagger

```go
http://127.0.0.1:9999/swagger/index.html

// 生成文档
$ go get -u github.com/swaggo/swag/cmd/swag
$ swag init
```

#### pprof 

```go
http://127.0.0.1:9999/debug/pprof
```

说明文档：

```go
// 查看 CPU 信息
go tool pprof 127.0.0.1:9999/debug/pprof/profile
...
(pprof) 

//输入 web，生成 svg 文件。
//输入 png，生成 png 文件。
//输入 top，查看排名前 20 的信息。
//查看更多命令，请执行 pprof help。
```

其他同理，比如：

```go
// 查看 内存 信息
go tool pprof 127.0.0.1:9999/debug/pprof/heap

// 查看 协程 信息
go tool pprof 127.0.0.1:9999/debug/pprof/goroutine

// 查看 锁 信息
go tool pprof 127.0.0.1:9999/debug/pprof/mutex
```
如果还想查看火焰图，请执行如下命令：

```go
// 1.下载 pprof 工具
go get -u github.com/google/pprof

// 2.启动可视化界面
pprof -http=:9998 xxx.cpu.prof

// 3.查看可视化界面
http://127.0.0.1:9998/ui/
```

## Special Thanks

[@koketama](https://github.com/koketama)

## Learning together

![](https://github.com/xinliangnote/Go/blob/master/00-基础语法/images/qr.jpg)

