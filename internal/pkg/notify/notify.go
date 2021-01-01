package notify

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/configs"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/notify/mail"

	"go.uber.org/zap"
)

// OnPanicNotify
func OnPanicNotify(ctx core.Context, err interface{}, stackInfo string) {
	cfg := configs.Get().Mail
	if cfg.Host == "" || cfg.Port == 0 || cfg.User == "" || cfg.Pass == "" || cfg.To == "" {
		ctx.Logger().Error("Mail config error")
		return
	}

	subject, body, htmlErr := mail.NewPanicHTMLEmail(ctx.Method(), ctx.Host(), ctx.URI(), ctx.Journal().ID(), err, stackInfo)
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
