package menu

import (
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
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
// @Security LoginToken
func (h *handler) DeleteAction() core.HandlerFunc {
	return func(c core.Context) {
		req := new(deleteActionRequest)
		res := new(deleteActionResponse)
		if err := c.ShouldBindURI(req); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}

		ids, err := h.hashids.HashidsDecode(req.Id)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.HashIdsDecodeError,
				code.Text(code.HashIdsDecodeError)).WithError(err),
			)
			return
		}

		id := int32(ids[0])

		err = h.menuService.DeleteAction(c, id)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.MenuDeleteActionError,
				code.Text(code.MenuDeleteActionError)).WithError(err),
			)
			return
		}

		res.Id = id
		c.Payload(res)
	}
}
