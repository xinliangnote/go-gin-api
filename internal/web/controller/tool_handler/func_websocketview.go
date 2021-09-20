package tool_handler

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

func (h *handler) WebsocketView() core.HandlerFunc {
	return func(c core.Context) {
		c.HTML("tool_websocket", nil)
	}
}
