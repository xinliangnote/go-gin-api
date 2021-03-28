package authorized_handler

import "github.com/xinliangnote/go-gin-api/internal/pkg/core"

func (h *handler) ListView() core.HandlerFunc {
	return func(c core.Context) {
		c.HTML("authorized_list", nil)
	}
}
