```
 ██████╗  ██████╗        ██████╗ ██╗███╗   ██╗       █████╗ ██████╗ ██╗
██╔════╝ ██╔═══██╗      ██╔════╝ ██║████╗  ██║      ██╔══██╗██╔══██╗██║
██║  ███╗██║   ██║█████╗██║  ███╗██║██╔██╗ ██║█████╗███████║██████╔╝██║
██║   ██║██║   ██║╚════╝██║   ██║██║██║╚██╗██║╚════╝██╔══██║██╔═══╝ ██║
╚██████╔╝╚██████╔╝      ╚██████╔╝██║██║ ╚████║      ██║  ██║██║     ██║
 ╚═════╝  ╚═════╝        ╚═════╝ ╚═╝╚═╝  ╚═══╝      ╚═╝  ╚═╝╚═╝     ╚═╝

* [register cors]
* [register rate]
* [register panic notify]
* [register pprof]
* [register swagger]
* [register prometheus]
```

## go-gin-api

基于 [Gin](https://github.com/gin-gonic/gin) 进行模块化设计的 API 框架，封装了常用的功能，使用简单，致力于进行快速的业务研发，同时增加了更多限制，以约束项目组开发成员、规避混乱无序和自由随意。

供参考学习，线上使用请谨慎！

**查看 Jaeger 链路追踪代码，请查看 [v1.0版](https://github.com/xinliangnote/go-gin-api/releases/tag/v1.0)，文档 [jaeger.md](https://github.com/xinliangnote/go-gin-api/blob/master/docs/jaeger.md) 点这里**。

持续更新... 

## Features

- [x] 包管理工具 [Go Modules](https://github.com/golang/go/wiki/Modules)
- [x] [Gin](https://github.com/gin-gonic/gin) 支持优雅关闭
- [x] 配置文件解析库 [Viper](https://github.com/spf13/viper)
- [x] 文档使用 [Swagger](https://swagger.io/) 生成
- [x] 性能分析使用 [pprof](github.com/gin-contrib/pprof)
- [x] 集成
    - [x] [JWT](https://jwt.io/) 身份认证
    - [x] [zap](go.uber.org/zap) 日志记录
    - [x] [rate](golang.org/x/time/rate) 限流
    - [x] 异常捕获并邮件告警
    - [x] 每个请求具备链路ID
    - [x] 统一定义错误码
    - [x] 支持 FAT、UAT、PRO 环境
    - [x] MD5、AES 对称加密、RSA 非对称加密
- [ ] 存储
    - [ ] MySQL
    - [ ] Redis
    - [ ] MongoDB
- [ ] gRPC
- [ ] ...

## Quick start

#### Requirements

- Go version >= 1.12
- Global environment configure (Linux/Mac)

```
export GO111MODULE=on
export GOPROXY=https://goproxy.io
```

#### Build & Run

```
cd go-gin-api

go run main.go
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

