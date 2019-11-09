package error

import (
	"fmt"
	"github.com/xinliangnote/go-util/json"
	"github.com/xinliangnote/go-util/mail"
	timeUtil "github.com/xinliangnote/go-util/time"
	"go-gin-api/app/config"
	"go-gin-api/app/route/middleware/exception"
	"log"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func ErrorNew (text string) error {
	alarm("INFO", text)
	return &errorString{text}
}

// 发邮件
func ErrorMail (text string) error {
	alarm("MAIL",text)
	return &errorString{text}
}

// 发短信
func ErrorSms (text string) error {
	alarm("SMS", text)
	return &errorString{text}
}

// 发微信
func ErrorWeChat (text string) error {
	alarm("WX", text)
	return &errorString{text}
}

// 告警方法
func  alarm(level string, str string) {
	if level == "MAIL" {
		DebugStack := ""
		for _, v := range strings.Split(string(debug.Stack()), "\n") {
			DebugStack += v + "<br>"
		}

		subject := fmt.Sprintf("【系统告警】%s 项目出错了！", config.AppName)

		body := strings.ReplaceAll(exception.MailTemplate, "{ErrorMsg}", fmt.Sprintf("%s", str))
		body  = strings.ReplaceAll(body, "{RequestTime}", timeUtil.GetCurrentDate())
		body  = strings.ReplaceAll(body, "{RequestURL}", "--")
		body  = strings.ReplaceAll(body, "{RequestUA}", "--")
		body  = strings.ReplaceAll(body, "{RequestIP}", "--")
		body  = strings.ReplaceAll(body, "{DebugStack}", DebugStack)

		// 执行发邮件
		options := &mail.Options{
			MailHost : config.SystemEmailHost,
			MailPort : config.SystemEmailPort,
			MailUser : config.SystemEmailUser,
			MailPass : config.SystemEmailPass,
			MailTo   : config.ErrorNotifyUser,
			Subject  : subject,
			Body     : body,
		}
		_ = mail.Send(options)

	} else if level == "SMS" {
		// 执行发短信

	} else if level == "WX" {
		// 执行发微信

	} else if level == "INFO" {
		// 执行记日志

		if f, err := os.OpenFile(config.AppErrorLogName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666); err != nil {
			log.Println(err)
		} else {
			errorLogMap := make(map[string]interface{})
			errorLogMap["time"] = time.Now().Format("2006/01/02 - 15:04:05")
			errorLogMap["info"] = str

			errorLogJson, _ := json.Encode(errorLogMap)
			_, _ = f.WriteString(errorLogJson + "\n")
		}
	}
}
