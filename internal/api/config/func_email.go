package config

import (
	"fmt"
	"net/http"

	"github.com/xinliangnote/go-gin-api/configs"
	"github.com/xinliangnote/go-gin-api/internal/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/pkg/env"
	"github.com/xinliangnote/go-gin-api/pkg/mail"

	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

type emailRequest struct {
	Host string `form:"host"` // 邮箱服务器
	Port string `form:"port"` // 端口
	User string `form:"user"` // 发件人邮箱
	Pass string `form:"pass"` // 发件人密码
	To   string `form:"to"`   // 收件人邮箱地址，多个用,分割
}

type emailResponse struct {
	Email string `json:"email"` // 邮箱地址
}

// Email 修改邮件配置
// @Summary 修改邮件配置
// @Description 修改邮件配置
// @Tags API.config
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param host formData string true "邮箱服务器"
// @Param port formData string true "端口"
// @Param user formData string true "发件人邮箱"
// @Param pass formData string true "发件人密码"
// @Param to formData string true "收件人邮箱地址，多个用,分割"
// @Success 200 {object} emailResponse
// @Failure 400 {object} code.Failure
// @Router /api/config/email [patch]
// @Security LoginToken
func (h *handler) Email() core.HandlerFunc {
	return func(c core.Context) {
		req := new(emailRequest)
		res := new(emailResponse)
		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}

		options := &mail.Options{
			MailHost: req.Host,
			MailPort: cast.ToInt(req.Port),
			MailUser: req.User,
			MailPass: req.Pass,
			MailTo:   req.To,
			Subject:  fmt.Sprintf("%s[%s] 邮箱告警人调整通知。", configs.ProjectName, env.Active().Value()),
			Body:     fmt.Sprintf("%s[%s] 已添加您为系统告警通知人。", configs.ProjectName, env.Active().Value()),
		}
		if err := mail.Send(options); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.SendEmailError,
				code.Text(code.SendEmailError)+err.Error()).WithError(err),
			)
			return
		}

		viper.Set("mail.host", req.Host)
		viper.Set("mail.port", cast.ToInt(req.Port))
		viper.Set("mail.user", req.User)
		viper.Set("mail.pass", req.Pass)
		viper.Set("mail.to", req.To)

		err := viper.WriteConfig()
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.WriteConfigError,
				code.Text(code.WriteConfigError)).WithError(err),
			)
			return
		}

		res.Email = req.To
		c.Payload(res)
	}
}
