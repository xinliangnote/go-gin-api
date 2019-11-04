package exception

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-api/app/config"
	"go-gin-api/app/util/mail"
	"go-gin-api/app/util/response"
	"go-gin-api/app/util/time"
	"runtime/debug"
	"strings"
)

func SetUp() gin.HandlerFunc {

	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {

				DebugStack := ""
				for _, v := range strings.Split(string(debug.Stack()), "\n") {
					DebugStack += v + "<br>"
				}

				subject := fmt.Sprintf("【重要错误】%s 项目出错了！", config.AppName)

				body := strings.ReplaceAll(MailTemplate, "{ErrorMsg}", fmt.Sprintf("%s", err))
				body  = strings.ReplaceAll(body, "{RequestTime}", time.GetCurrentDate())
				body  = strings.ReplaceAll(body, "{RequestURL}", c.Request.Method + "  " + c.Request.Host + c.Request.RequestURI)
				body  = strings.ReplaceAll(body, "{RequestUA}", c.Request.UserAgent())
				body  = strings.ReplaceAll(body, "{RequestIP}", c.ClientIP())
				body  = strings.ReplaceAll(body, "{DebugStack}", DebugStack)

				_ = mail.SendMail(config.ErrorNotifyUser, subject, body)

				utilGin := response.Gin{Ctx: c}
				utilGin.Response(500, "系统异常，请联系管理员！", nil)
			}
		}()
		c.Next()
	}
}
