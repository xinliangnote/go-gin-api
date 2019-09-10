package requestid

import (
	"github.com/gin-gonic/gin"
	"go-gin-api/app/util"
)

func SetUp() gin.HandlerFunc {

	return func(c *gin.Context) {
		requestId := c.Request.Header.Get("X-Request-Id")
		if requestId == "" {
			requestId = util.GenUUID()
		}
		c.Set("X-Request-Id", requestId)
		c.Writer.Header().Set("X-Request-Id", requestId)
		c.Next()
	}
}
