package route

import (
	"github.com/gin-gonic/gin"
	"go-gin-api/app/controller/product"
	"go-gin-api/app/route/middleware/logger"
	"go-gin-api/app/util"
)

func SetupRouter(engine *gin.Engine) {

	engine.Use(logger.SetUp())

	//404
	engine.NoRoute(func(c *gin.Context) {
		utilGin := util.Gin{Ctx:c}
		utilGin.Response(404,"请求方法不存在", nil)
	})

	engine.GET("/ping", func(c *gin.Context) {
		utilGin := util.Gin{Ctx:c}
		utilGin.Response(1,"pong", nil)
	})

	//@todo 记录请求超时的路由

	ProductRouter := engine.Group("/product")
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
