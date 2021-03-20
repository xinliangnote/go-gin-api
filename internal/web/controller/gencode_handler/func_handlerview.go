package gencode_handler

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

func (h *handler) HandlerView() core.HandlerFunc {
	return func(c core.Context) {
		c.HTML("gencode_handler", nil)
	}
}
