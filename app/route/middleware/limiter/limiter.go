package limiter

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-api/app/util/response"
	"golang.org/x/time/rate"
	"time"
)

func SetUp (maxBurstSize int) gin.HandlerFunc {

	limiter := rate.NewLimiter(rate.Every(time.Second*1), maxBurstSize)
	return func(c *gin.Context) {
		if limiter.Allow() {
			c.Next()
			return
		}
		fmt.Println("Too many requests")
		utilGin := response.Gin{Ctx: c}
		utilGin.Response(-1, "Too many requests", nil)
		c.Abort()
		return
	}
}
