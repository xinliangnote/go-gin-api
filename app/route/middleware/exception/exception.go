package exception

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-api/app/config"
	"go-gin-api/app/util"
	"runtime/debug"
	"strings"
	"time"
)

func SetUp() gin.HandlerFunc {

	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				subject := fmt.Sprintf("[Panic - %s] 项目出错了！", config.AppName)
				body := fmt.Sprintf("<b>错误时间：%s\n Runtime：\n</b>%s",time.Now().Format("2006/01/02 - 15:04:05"), string(debug.Stack()))
				bodyHtml := ""
				for _,v := range strings.Split(body, "\n") {
					bodyHtml += v + "<br>"
				}
				_ = util.SendMail(config.ErrorNotifyUser, subject, bodyHtml)

				utilGin := util.Gin{Ctx:c}
				utilGin.Response(500, "系统异常，请联系管理员！", nil)
			}
		}()
		c.Next()
	}
}
