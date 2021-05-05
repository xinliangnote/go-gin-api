package admin_handler

import "github.com/xinliangnote/go-gin-api/internal/pkg/core"

func (h *handler) MenuView() core.HandlerFunc {
	return func(c core.Context) {
		c.HTML("menu_view", nil)
	}
}
