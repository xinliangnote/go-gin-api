package menu_handler

import (
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/api/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/pkg/errno"
)

type deleteActionRequest struct {
	Id string `uri:"id"` // HashID
}

type deleteActionResponse struct {
	Id int32 `json:"id"` // 主键ID
}

// DeleteAction 删除功能权限
// @Summary 删除功能权限
// @Description 删除功能权限
// @Tags API.menu
// @Accept json
// @Produce json
// @Param id path string true "hashId"
// @Success 200 {object} deleteActionResponse
// @Failure 400 {object} code.Failure
// @Router /api/menu_action/{id} [delete]
func (h *handler) DeleteAction() core.HandlerFunc {
	return func(c core.Context) {
		req := new(deleteActionRequest)
		res := new(deleteActionResponse)
		if err := c.ShouldBindURI(req); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithErr(err),
			)
			return
		}

		ids, err := h.hashids.HashidsDecode(req.Id)
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.HashIdsDecodeError,
				code.Text(code.HashIdsDecodeError)).WithErr(err),
			)
			return
		}

		id := int32(ids[0])

		err = h.menuService.DeleteAction(c, id)
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.MenuDeleteActionError,
				code.Text(code.MenuDeleteActionError)).WithErr(err),
			)
			return
		}

		res.Id = id
		c.Payload(res)
	}
}
