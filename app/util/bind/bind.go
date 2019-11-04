package bind

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func Bind(s interface{}, c *gin.Context) (interface{}, error) {
	b := binding.Default(c.Request.Method, c.ContentType())
	if err := c.ShouldBindWith(s, b); err != nil {
		return nil, err
	}
	return s, nil
}
