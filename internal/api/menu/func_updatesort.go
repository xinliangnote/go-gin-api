package menu

import (
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

type updateSortRequest struct {
	Id   string `form:"id"`   // HashId
	Sort int32  `form:"sort"` // 排序
}

type updateSortResponse struct {
	Id int32 `json:"id"` // 主键ID
}

// UpdateSort 更新菜单排序
// @Summary 更新菜单排序
// @Description 更新菜单排序
// @Tags API.menu
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param id formData string true "hashId"
// @Param sort formData int true "排序"
// @Success 200 {object} updateSortResponse
// @Failure 400 {object} code.Failure
// @Router /api/menu/sort [patch]
// @Security LoginToken
func (h *handler) UpdateSort() core.HandlerFunc {
	return func(c core.Context) {
		req := new(updateSortRequest)
		res := new(updateSortResponse)
		if err := c.ShouldBindForm(req); err != nil {
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

		err = h.menuService.UpdateSort(c, id, req.Sort)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.MenuUpdateError,
				code.Text(code.MenuUpdateError)).WithError(err),
			)
			return
		}

		res.Id = id
		c.Payload(res)
	}
}
