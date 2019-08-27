package param_bind

import (
	"github.com/gin-gonic/gin/binding"
	"go-gin-api/app/controller/param_verify"
	"gopkg.in/go-playground/validator.v8"
)

type ProductAdd struct {
	Name string `form:"name" json:"name" binding:"required,NameValid"`
}

func init() {
	// 绑定验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("NameValid", param_verify.NameValid)
	}
}
