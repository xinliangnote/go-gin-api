![](https://github.com/xinliangnote/Go/blob/master/03-go-gin-api%20%5B%E6%96%87%E6%A1%A3%5D/images/go-gin-api-logo.png)

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

## Download

```
git clone https://github.com/xinliangnote/go-gin-api.git
```

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

输出如下，表示 Http Server 启动成功。
|-----------------------------------|
|            go-gin-api             |
|-----------------------------------|
|  Go Http Server Start Successful  |
|    Port:9999     Pid:xxxxx        |
|-----------------------------------|
```

#### HTTP Demo

```
curl -X POST http://127.0.0.1:9999/product
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

## Dependence

- WEB 框架：github.com/gin-gonic/gin
- 链路追踪：github.com/jaegertracing/jaeger-client-go
- 限流：golang.org/x/time/rate
- 工具包：github.com/xinliangnote/go-util

## Document

- [1. 使用 go modules 初始化项目](https://mp.weixin.qq.com/s/1XNTEgZ0XGZZdxFOfR5f_A)
- [2. 规划项目目录和参数验证](https://mp.weixin.qq.com/s/11AuXptWGmL5QfiJArNLnA)
- [3. 路由中间件 - 日志记录](https://mp.weixin.qq.com/s/eTygPXnrYM2xfrRQyfn8Tg)
- [4. 路由中间件 - 捕获异常](https://mp.weixin.qq.com/s/SconDXB_x7Gan6T0Awdh9A)
- [5. 路由中间件 - Jaeger 链路追踪（理论篇）](https://mp.weixin.qq.com/s/28UBEsLOAHDv530ePilKQA)
- [6. 路由中间件 - Jaeger 链路追踪（实战篇）](https://mp.weixin.qq.com/s/Ea28475_UTNaM9RNfgPqJA)
- [7. 路由中间件 - 签名验证](https://mp.weixin.qq.com/s/0cozELotcpX3Gd6WPJiBbQ)

## Learning together

![](https://github.com/xinliangnote/Go/blob/master/00-基础语法/images/qr.jpg)

