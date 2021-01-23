package notify

import (
	"github.com/xinliangnote/go-gin-api/configs"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/pkg/mail"

	"go.uber.org/zap"
)

// OnPanicNotify 发生 panic 时进行通知
func OnPanicNotify(ctx core.Context, err interface{}, stackInfo string) {
	cfg := configs.Get().Mail
	if cfg.Host == "" || cfg.Port == 0 || cfg.User == "" || cfg.Pass == "" || cfg.To == "" {
		ctx.Logger().Error("Mail config error")
		return
	}

	subject, body, htmlErr := NewPanicHTMLEmail(
		ctx.Method(),
		ctx.Host(),
		ctx.URI(),
		ctx.Trace().ID(),
		err,
		stackInfo,
	)
	if htmlErr != nil {
		ctx.Logger().Error("NewPanicHTMLEmail error", zap.Error(htmlErr))
		return
	}

	options := &mail.Options{
		MailHost: cfg.Host,
		MailPort: cfg.Port,
		MailUser: cfg.User,
		MailPass: cfg.Pass,
		MailTo:   cfg.To,
		Subject:  subject,
		Body:     body,
	}
	sendErr := mail.Send(options)
	if sendErr != nil {
		ctx.Logger().Error("Mail Send error", zap.Error(sendErr))
	}

	return
}
