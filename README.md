![](https://github.com/xinliangnote/Go/blob/master/03-go-gin-api%20%5B%E6%96%87%E6%A1%A3%5D/images/go-gin-api-logo.png)

## go-gin-api

基于 [Gin](https://github.com/gin-gonic/gin) 进行模块化设计的 API 框架，封装了常用的功能，使用简单，致力于进行快速的业务研发。

持续更新... 

## Features

- [x] 使用 go modules 初始化项目
- [x] 安装 Gin 框架
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
    - [ ] 限流
- [ ] 自定义告警
    - [ ] 邮件（gomail）
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

#### Jaeger 链路追踪 Demo

```
http://127.0.0.1:9999/jaeger_test
```

## Jaeger 链路追踪效果

![](https://github.com/xinliangnote/Go/blob/master/03-go-gin-api%20%5B文档%5D/images/jaeger_demo_2.png)

![](https://github.com/xinliangnote/Go/blob/master/03-go-gin-api%20%5B文档%5D/images/jaeger_demo_3.png)

## Documents

- [帮助文档](https://github.com/xinliangnote/Go/tree/master/03-go-gin-api%20%5B文档%5D/)
- [Jaeger 链路追踪 Service 端](https://github.com/xinliangnote/go-jaeger-demo)

## Learning together

![](https://github.com/xinliangnote/Go/blob/master/00-基础语法/images/qr.jpg)

