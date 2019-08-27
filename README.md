![](https%3a%2f%2fgithub.com%2fxinliangnote%2fGo%2fblob%2fmaster%2f03-go-gin-api+%5b%e6%96%87%e6%a1%a3%5d%2fimages%2fgo-gin-api-logo.png)

![](https://camo.githubusercontent.com/0f4b285c5516603a1c2a085dbfc67a6505155edb/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f4c616e67756167652d476f2d626c75652e737667)

## go-gin-api

[Go 项目实战] 实现一个开箱即用的 API 框架的轮子，这个轮子是基于 Gin 框架的基础上开发的。

持续更新... 

## Features

-[x] 使用 go modules 初始化项目
-[x] 安装 Gin 框架
-[x] 规划项目目录
-[x] 参数验证
    -[x] 模型绑定和验证
    -[x] 自定义验证器
-[ ] 路由中间件
    -[ ] 签名验证
        -[ ] MD5 组合拳
        -[ ] AES 对称加密
        -[ ] RSA 非对称加密
    -[ ] 日志记录
    -[ ] 异常捕获
    -[ ] Jaeger 链路追踪
-[ ] 自定义告警
    -[ ] 邮件
    -[ ] 微信
    -[ ] 短信
    -[ ] 钉钉
-[ ] gRPC
-[ ] ...

## Quick start

#### Requirements

- Go version >= 1.12
- Global environment configure (Linux/Mac)

```
export GO111MODULE=on
export GOPROXY=https://goproxy.io
```

#### Installation

```
go get github.com/xinliangnote/go-gin-api
```

#### Build & Run

```
cd go-gin-api

go run main.go
```

#### Test demo

```
curl -X POST http://127.0.0.1:9999/product
```

## Documents

- [中文文档](https://github.com/xinliangnote/Go/tree/master/03-go-gin-api%20%5B文档%5D/)

## Learning together

![](https://github.com/xinliangnote/Go/blob/master/00-基础语法/images/qr.jpg)

