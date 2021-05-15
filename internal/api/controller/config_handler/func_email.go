package config_handler

import (
	"fmt"
	"net/http"

	"github.com/xinliangnote/go-gin-api/configs"
	"github.com/xinliangnote/go-gin-api/internal/api/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/pkg/env"
	"github.com/xinliangnote/go-gin-api/pkg/errno"
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
// @Accept multipart/form-data
// @Produce json
// @Param host formData string true "邮箱服务器"
// @Param port formData string true "端口"
// @Param user formData string true "发件人邮箱"
// @Param pass formData string true "发件人密码"
// @Param to formData string true "收件人邮箱地址，多个用,分割"
// @Success 200 {object} emailResponse
// @Failure 400 {object} code.Failure
// @Router /api/config/email [patch]
func (h *handler) Email() core.HandlerFunc {
	return func(c core.Context) {
		req := new(emailRequest)
		res := new(emailResponse)
		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithErr(err),
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
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ConfigEmailError,
				"Mail Send error: "+err.Error()).WithErr(err),
			)
			return
		}

		viper.SetConfigName(env.Active().Value() + "_configs")
		viper.SetConfigType("toml")
		viper.AddConfigPath("./configs")

		viper.Set("mail.host", req.Host)
		viper.Set("mail.port", cast.ToInt(req.Port))
		viper.Set("mail.user", req.User)
		viper.Set("mail.pass", req.Pass)
		viper.Set("mail.to", req.To)

		err := viper.WriteConfig()
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ConfigEmailError,
				code.Text(code.ConfigEmailError)).WithErr(err),
			)
			return
		}

		res.Email = req.To
		c.Payload(res)
	}
}
