## go-gin-api ![go report](https://goreportcard.com/badge/github.com/xinliangnote/go-gin-api)

[![GitHub stars](https://img.shields.io/github/stars/xinliangnote/go-gin-api)](https://github.com/xinliangnote/go-gin-api/stargazers)
[![GitHub forks](https://img.shields.io/github/forks/xinliangnote/go-gin-api)](https://github.com/xinliangnote/go-gin-api/network/members)
[![GitHub watchers](https://img.shields.io/github/watchers/xinliangnote/go-gin-api)](https://github.com/xinliangnote/go-gin-api/watchers)
[![GitHub license](https://img.shields.io/github/license/xinliangnote/go-gin-api)](https://github.com/xinliangnote/go-gin-api/blob/master/LICENSE)
[![GitHub repo size](https://img.shields.io/github/repo-size/xinliangnote/go-gin-api)](https://github.com/xinliangnote/go-gin-api)

基于 [Gin](https://github.com/gin-gonic/gin) 进行模块化设计的 API 框架，封装了常用的功能，使用简单，致力于进行快速的业务研发，同时增加了更多限制，以约束项目组开发成员、规避混乱无序和自由随意。

查看 Jaeger 链路追踪代码，请查看 [v1.0版](https://github.com/xinliangnote/go-gin-api/releases/tag/v1.0)，文档点这里 [jaeger.md](https://github.com/xinliangnote/go-gin-api/blob/master/docs/jaeger.md) 。

供参考学习，线上使用请谨慎！

持续更新... 

## Star Module

亮点功能：:star: [trace] 开发调试的辅助工具。

可记录如下信息：

- trace_id，当前请求的唯一ID
- request，当前请求的请求信息
- response，当前请求的返回信息
- third_party_requests，当前请求涉及到调用第三方的信息
- debugs，当前请求的调试信息
- sqls，当前请求执行的 sql 信息
- redis，当前请求执行的 redis 信息
- success，当前请求结果
- cost_seconds，执行时长，单位秒

## Catalogue

```cassandraql
├── main.go                      # 项目入口文件
├── configs                      # 配置文件统一存放目录
├── docs                         # Swagger 文档，执行 swag init 生成的
├── init                         # 项目初始脚本，比如初始化 SQL 语句
├── internal                     # 业务目录
│   ├── api                      # 业务代码
│   ├── pkg                      # 内部使用的 package
├── logs                         # 存放日志的目录
└── pkg                          # 一些封装好的 package
```
## Features

- [x] 包管理工具 [Go Modules](https://github.com/golang/go/wiki/Modules)
- [x] [Gin](https://github.com/gin-gonic/gin) 支持优雅关闭
- [x] 配置文件解析库 [Viper](https://github.com/spf13/viper)
- [x] 文档使用 [Swagger](https://swagger.io/) 生成
- [x] 性能分析使用 [pprof](https://github.com/gin-contrib/pprof)
- [x] [zap](https://go.uber.org/zap) 日志记录
- [x] [rate](https://golang.org/x/time/rate) 限流
- [x] [token] 基于[JWT](https://github.com/dgrijalva/jwt-go) 身份认证
- [x] [notify] 异常捕获并进行邮件告警
- [x] [trace] 开发调试的辅助工具
    - [x] 支持设置 trace_id
    - [x] 支持设置 request 信息
    - [x] 支持设置 response 信息
    - [x] 支持设置 third_party_requests 三方请求信息
    - [x] 支持设置 debugs 打印调试信息
    - [x] 支持设置 sqls 执行 SQL 信息
    - [x] 支持设置 redis 执行 Redis 信息
    - [x] 可记录 cost_seconds 执行时长
- [x] [errno] 统一定义错误码
- [x] [env] 支持 FAT、UAT、PRO 环境
- [x] [aes] AES 对称加密
    - [x] 支持密码分组链模式（CBC）
- [x] [rsa] RSA 非对称加密
- [x] 数据库组件使用 [GORM V2](https://gorm.io/gorm)
    - [x] 自带连接池
- [x] Redis 组件使用 [go-redis](https://github.com/go-redis/redis)
    - [x] 自带连接池
- [ ] MongoDB
- [x] 使用 [Prometheus Client](https://github.com/prometheus/client_golang/prometheus)
    - [x] 已设置计数器（Counter）指标
    - [x] 已设置直方图（Histogram）指标
- [ ] 任务调度
- [ ] gRPC
- [ ] ID 生成器
- [x] [httpclient] HTTP 请求包
    - [x] 支持设置 TTL 一次请求的最长执行时间
    - [x] 支持设置 Header 信息
    - [x] 支持设置 Trace 信息
    - [x] 支持设置 Logger 信息
    - [x] 支持设置 Mock 数据
    - [x] 支持设置 OnFailedAlarm 失败
        - [x] 可设置 alarmTitle 告警标题
        - [x] 可设置 alarmObject 告警方式（邮件/短信/微信）
        - [x] 可设置 alarmVerify 定义符合告警的验证规则
    - [x] 支持设置 OnFailedRetry 失败重试
        - [x] 可设置 retryTimes 重试次数
        - [x] 可设置 retryDelay 重试前延迟等待时间
        - [x] 可设置 retryVerify 定义符合重试的验证规则

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

go run main.go -env fat

 ██████╗  ██████╗        ██████╗ ██╗███╗   ██╗       █████╗ ██████╗ ██╗
██╔════╝ ██╔═══██╗      ██╔════╝ ██║████╗  ██║      ██╔══██╗██╔══██╗██║
██║  ███╗██║   ██║█████╗██║  ███╗██║██╔██╗ ██║█████╗███████║██████╔╝██║
██║   ██║██║   ██║╚════╝██║   ██║██║██║╚██╗██║╚════╝██╔══██║██╔═══╝ ██║
╚██████╔╝╚██████╔╝      ╚██████╔╝██║██║ ╚████║      ██║  ██║██║     ██║
 ╚═════╝  ╚═════╝        ╚═════╝ ╚═╝╚═╝  ╚═══╝      ╚═╝  ╚═╝╚═╝     ╚═╝

* [register port :9999]
* [register env fat]
* [register cors]
* [register rate]
* [register panic notify]
* [register pprof]
* [register swagger]
* [register prometheus]

// 启动成功后，可访问：http://127.0.0.1:9999/system/health
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

