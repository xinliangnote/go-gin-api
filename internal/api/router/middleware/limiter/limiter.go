package limiter

import (
	"net/http"
	"time"

	"github.com/xinliangnote/go-gin-api/internal/pkg/errno"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func SetUp(maxBurstSize int) gin.HandlerFunc {

	limiter := rate.NewLimiter(rate.Every(time.Second*1), maxBurstSize)
	return func(c *gin.Context) {
		if limiter.Allow() {
			c.Next()
			return
		}
		c.JSON(http.StatusOK, errno.ErrManyRequest)
		c.Abort()
		return
	}
}
