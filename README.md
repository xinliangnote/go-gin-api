## 关于

`go-gin-api` 是基于 [Gin](https://github.com/gin-gonic/gin) 进行模块化设计的 API 框架，封装了常用的功能，使用简单，致力于进行快速的业务研发，同时增加了更多限制，约束项目组开发成员，规避混乱无序及自由随意的编码。

供参考学习，线上使用请谨慎！

集成组件：

1. 支持 [rate](https://golang.org/x/time/rate) 接口限流 
1. 支持 panic 异常时邮件通知 
1. 支持 [cors](https://github.com/rs/cors) 接口跨域 
1. 支持 [Prometheus](https://github.com/prometheus/client_golang) 指标记录 
1. 支持 [Swagger](https://github.com/swaggo/gin-swagger) 接口文档生成 
1. 支持 [GraphQL](https://github.com/99designs/gqlgen) 查询语言 
1. 支持 trace 项目内部链路追踪 
1. 支持 [pprof](https://github.com/gin-contrib/pprof) 性能剖析
1. 支持 errno 统一定义错误码 
1. 支持 [zap](https://go.uber.org/zap) 日志收集 
1. 支持 [viper](https://github.com/spf13/viper) 配置文件解析
1. 支持 [gorm](https://gorm.io/gorm) 数据库组件
1. 支持 [go-redis](https://github.com/go-redis/redis/v7) 组件
1. 支持 RESTful API 返回值规范
1. 支持 生成数据表 CURD、控制器方法 等代码生成器
1. 支持 [cron](https://github.com/jakecoffman/cron) 定时任务，在后台可界面配置
1. 支持 [websocket](https://github.com/gorilla/websocket) 实时通讯，在后台有界面演示
1. 支持 web 界面，使用的 [Light Year Admin 模板](https://gitee.com/yinqi/Light-Year-Admin-Using-Iframe)


## 文档索引（可加入交流群）

- 中文文档：[go-gin-api - 语雀](https://www.yuque.com/xinliangnote/go-gin-api/ngc3x5)
- English Document：[en.md](https://github.com/xinliangnote/go-gin-api/blob/master/en.md)

## 轻量版

为了满足开发者对于简单、轻量级 API 框架的需求，开发了 gin-api-mono，旨在提供更便捷的业务开发体验。

相比于 go-gin-api，首先 gin-api-mono 去掉了一些集成的功能和界面，使得整个框架更加简洁、轻量。其次 gin-api-mono 对框架代码进行了升级，以确保其在性能和稳定性方面的优势。这样，开发者就可以更灵活地选择所需的功能，并获得更好的性能和稳定性。

详见链接：https://xiaobot.net/post/e9f7ef4c-81b1-4ffc-9053-bec55c3abb12

## 其他

查看 Jaeger 链路追踪 Demo 代码，请查看 [v1.0 版](https://github.com/xinliangnote/go-gin-api/releases/tag/v1.0) ，链接地址：http://127.0.0.1:9999/jaeger_test

调用的其他服务端 Demo 代码为 [https://github.com/xinliangnote/go-jaeger-demo](https://github.com/xinliangnote/go-jaeger-demo)

## 联系作者

![联系作者](https://i.loli.net/2021/07/02/cwiLQ13CRgJIS86.jpg)

