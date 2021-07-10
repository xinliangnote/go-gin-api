package admin_handler

import (
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/pkg/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/pkg/errno"
)

type menuActionViewRequest struct {
	Id string `uri:"id"` // 主键ID
}

type menuActionViewResponse struct {
	HashID string `json:"hash_id"` // hashID
}

func (h *handler) MenuActionView() core.HandlerFunc {
	return func(c core.Context) {
		req := new(menuActionViewRequest)
		if err := c.ShouldBindURI(req); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithErr(err),
			)
			return
		}

		obj := new(menuActionViewResponse)
		obj.HashID = req.Id

		c.HTML("menu_action", obj)
	}
}
