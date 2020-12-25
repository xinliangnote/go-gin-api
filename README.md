```
 ██████╗  ██████╗        ██████╗ ██╗███╗   ██╗       █████╗ ██████╗ ██╗
██╔════╝ ██╔═══██╗      ██╔════╝ ██║████╗  ██║      ██╔══██╗██╔══██╗██║
██║  ███╗██║   ██║█████╗██║  ███╗██║██╔██╗ ██║█████╗███████║██████╔╝██║
██║   ██║██║   ██║╚════╝██║   ██║██║██║╚██╗██║╚════╝██╔══██║██╔═══╝ ██║
╚██████╔╝╚██████╔╝      ╚██████╔╝██║██║ ╚████║      ██║  ██║██║     ██║
 ╚═════╝  ╚═════╝        ╚═════╝ ╚═╝╚═╝  ╚═══╝      ╚═╝  ╚═╝╚═╝     ╚═╝
```

## go-gin-api

基于 [Gin](https://github.com/gin-gonic/gin) 进行模块化设计的 API 框架，封装了常用的功能，使用简单，致力于进行快速的业务研发。

持续更新... 

## Features

- [x] 使用 go modules 初始化项目
- [x] 安装 Gin 框架
- [x] 性能分析工具（pprof）
- [x] 支持优雅地重启或停止
- [x] 规划项目目录
- [x] 参数验证（validator.v9）
    - [x] 模型绑定和验证
    - [x] 自定义验证器
- [x] 路由中间件
    - [x] 签名验证
        - [x] MD5 组合加密
        - [x] AES 对称加密
        - [x] RSA 非对称加密
    - [x] 日志记录
    - [x] 异常捕获
    - [x] 链路追踪（Jaeger）
    - [x] 限流
- [x] 自定义告警
    - [x] 邮件（gomail）
    - [ ] 微信
    - [ ] 短信
    - [ ] 钉钉
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

#### HTTP Demo

```
curl -X POST http://127.0.0.1:9999/demo/user
```

#### Jaeger Demo

访问：

```
http://127.0.0.1:9999/jaeger_test
```

服务端测试代码：

- [https://github.com/xinliangnote/go-jaeger-demo](https://github.com/xinliangnote/go-jaeger-demo)

![](https://github.com/xinliangnote/Go/blob/master/03-go-gin-api%20%5B文档%5D/images/jaeger_demo_2.png)

![](https://github.com/xinliangnote/Go/blob/master/03-go-gin-api%20%5B文档%5D/images/jaeger_demo_3.png)

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

