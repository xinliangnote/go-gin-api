package product

import (
	"github.com/gin-gonic/gin"
	"go-gin-api/app/controller/param_bind"
	"go-gin-api/app/util"
)

// 新增
func Add(c *gin.Context) {
	utilGin := util.Gin{Ctx:c}
	if err := c.ShouldBind(&param_bind.ProductAdd{}); err != nil {
		utilGin.Response(-1, err.Error(), nil)
		return
	}

	// 业务处理...

	utilGin.Response(1, "success", nil)
}

// 编辑
func Edit(c *gin.Context) {

}

// 删除
func Delete(c *gin.Context) {

}

// 详情

func Detail(c *gin.Context) {

}
