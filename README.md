![](https://github.com/xinliangnote/Go/blob/master/03-go-gin-api%20%5B%E6%96%87%E6%A1%A3%5D/images/go-gin-api-logo.png)

## go-gin-api

基于 Gin 进行模块化设计的 API 框架，封装了常用的功能，使用简单，致力于进行快速的业务研发。

持续更新... 

## Features

- [x] 使用 go modules 初始化项目
- [x] 安装 Gin 框架
- [x] 优雅地重启或停止
- [x] 规划项目目录
- [x] 参数验证
    - [x] 模型绑定和验证
    - [x] 自定义验证器
- [ ] 路由中间件
    - [ ] 签名验证
        - [ ] MD5 组合拳
        - [ ] AES 对称加密
        - [ ] RSA 非对称加密
    - [ ] 日志记录
    - [ ] 异常捕获
    - [ ] Jaeger 链路追踪
- [ ] 自定义告警
    - [ ] 邮件
    - [ ] 微信
    - [ ] 短信
    - [ ] 钉钉
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
```

#### Test demo

```
curl -X POST http://127.0.0.1:9999/product
```

## Documents

- [中文文档](https://github.com/xinliangnote/Go/tree/master/03-go-gin-api%20%5B文档%5D/)

## Learning together

![](https://github.com/xinliangnote/Go/blob/master/00-基础语法/images/qr.jpg)

