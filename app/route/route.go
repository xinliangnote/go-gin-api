package route

import (
	"github.com/gin-gonic/gin"
	"go-gin-api/app/controller/product"
	"go-gin-api/app/util"
)

func SetupRouter(engine *gin.Engine) {
	engine.GET("/ping", func(c *gin.Context) {
		utilGin := util.Gin{Ctx:c}
		utilGin.Response(1,"pong", nil)
	})

	ProductRouter := engine.Group("")
	{
		// 新增产品
		ProductRouter.POST("/product", product.Add)

		// 更新产品
		ProductRouter.PUT("/product/:id", product.Edit)

		// 删除产品
		ProductRouter.DELETE("/product/:id", product.Delete)

		// 获取产品详情
		ProductRouter.GET("/product/:id", product.Detail)
	}
}
