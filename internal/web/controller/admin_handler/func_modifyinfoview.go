package admin_handler

import "github.com/xinliangnote/go-gin-api/internal/pkg/core"

func (h *handler) ModifyInfoView() core.HandlerFunc {
	return func(c core.Context) {
		c.HTML("admin_modifyinfo", nil)
	}
}
