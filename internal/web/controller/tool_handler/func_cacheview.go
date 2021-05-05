package tool_handler

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

func (h *handler) CacheView() core.HandlerFunc {
	return func(c core.Context) {
		c.HTML("tool_cache", nil)
	}
}
