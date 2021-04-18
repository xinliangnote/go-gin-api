package config_handler

import (
	"github.com/xinliangnote/go-gin-api/configs"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

func (h *handler) EmailView() core.HandlerFunc {
	return func(c core.Context) {
		c.HTML("config_email", configs.Get())
	}
}
