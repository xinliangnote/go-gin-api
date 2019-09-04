package exception

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-api/app/config"
	"go-gin-api/app/util"
	"runtime/debug"
	"strings"
)

func SetUp() gin.HandlerFunc {

	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				subject := fmt.Sprintf("【重要错误】%s 项目出错了！", config.AppName)

				body := fmt.Sprintf("<b>ErrorMessage: </b> %s \n", err)
				body += fmt.Sprintf("<b>RequestTime: </b> %s \n", util.GetCurrentDate())
				body += fmt.Sprintf("<b>RequestURL: </b> %s %s \n", c.Request.Method, c.Request.RequestURI)
				body += fmt.Sprintf("<b>RequestProto: </b> %s \n", c.Request.Proto)
				body += fmt.Sprintf("<b>RequestReferer: </b> %s \n", c.Request.Referer())
				body += fmt.Sprintf("<b>RequestUA: </b> %s \n", c.Request.UserAgent())
				body += fmt.Sprintf("<b>RequestClientIp: </b> %s \n", c.ClientIP())
				body += fmt.Sprintf("<b>DebugStack: </b> %s \n", string(debug.Stack()))

				bodyHtml := ""
				for _, v := range strings.Split(body, "\n") {
					bodyHtml += v + "<br>"
				}
				_ = util.SendMail(config.ErrorNotifyUser, subject, bodyHtml)

				utilGin := util.Gin{Ctx: c}
				utilGin.Response(500, "系统异常，请联系管理员！", nil)
			}
		}()
		c.Next()
	}
}
