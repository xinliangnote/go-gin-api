package tool_handler

import (
	"github.com/xinliangnote/go-gin-api/configs"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

func (h *handler) HashIdsView() core.HandlerFunc {
	return func(c core.Context) {
		c.HTML("tool_hashids", configs.Get())
	}
}
