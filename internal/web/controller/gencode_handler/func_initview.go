package gencode_handler

import "github.com/xinliangnote/go-gin-api/internal/pkg/core"

func (h *handler) InitView() core.HandlerFunc {
	return func(c core.Context) {
		c.HTML("gencode_init", nil)
	}
}
