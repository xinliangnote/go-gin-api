package mail

import (
	"fmt"
	"go-gin-api/app/config"
	"gopkg.in/gomail.v2"
	"strings"
)

func SendMail(mailTo string, subject string, body string) error {
	
	if config.ErrorNotifyOpen != 1 {
		return nil
	}

	m := gomail.NewMessage()

	//设置发件人
	m.SetHeader("From", config.SystemEmailUser)

	//设置发送给多个用户
	mailArrTo := strings.Split(mailTo, ",")
	m.SetHeader("To", mailArrTo...)

	//设置邮件主题
	m.SetHeader("Subject", subject)

	//设置邮件正文
	m.SetBody("text/html", body)

	d := gomail.NewDialer(config.SystemEmailHost, config.SystemEmailPort, config.SystemEmailUser, config.SystemEmailPass)

	err := d.DialAndSend(m)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
