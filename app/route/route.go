package route

import (
	"github.com/gin-gonic/gin"
	"go-gin-api/app/controller/jaeger_conn"
	"go-gin-api/app/controller/product"
	"go-gin-api/app/route/middleware/exception"
	"go-gin-api/app/route/middleware/jaeger"
	"go-gin-api/app/route/middleware/logger"
	signMD5 "go-gin-api/app/route/middleware/sign/md5"
	"go-gin-api/app/util"
)

func SetupRouter(engine *gin.Engine) {

	//设置路由中间件
	engine.Use(logger.SetUp(), exception.SetUp(), jaeger.SetUp())

	//404
	engine.NoRoute(func(c *gin.Context) {
		utilGin := util.Gin{Ctx:c}
		utilGin.Response(404,"请求方法不存在", nil)
	})

	engine.GET("/ping", func(c *gin.Context) {
		utilGin := util.Gin{Ctx:c}
		utilGin.Response(1,"pong", nil)
	})

	// 测试链路追踪
	engine.GET("/jaeger_test", jaeger_conn.JaegerTest)

	//@todo 记录请求超时的路由

	ProductRouter := engine.Group("/product").Use(signMD5.SetUp())
	{
		// 新增产品
		ProductRouter.POST("", product.Add)

		// 更新产品
		ProductRouter.PUT("/:id", product.Edit)

		// 删除产品
		ProductRouter.DELETE("/:id", product.Delete)

		// 获取产品详情
		ProductRouter.GET("/:id", product.Detail)
	}
}
