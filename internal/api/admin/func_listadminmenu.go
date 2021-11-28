package admin

import (
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/services/admin"
)

type listAdminMenuRequest struct {
	Id string `uri:"id"` // HashID
}

type listAdminMenuResponse struct {
	List     []admin.ListMenuData `json:"list"`
	UserName string               `json:"username"`
}

// ListAdminMenu 菜单授权列表
// @Summary 菜单授权列表
// @Description 菜单授权列表
// @Tags API.admin
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param id path string true "hashId"
// @Success 200 {object} listAdminMenuResponse
// @Failure 400 {object} code.Failure
// @Router /api/admin/menu/{id} [get]
// @Security LoginToken
func (h *handler) ListAdminMenu() core.HandlerFunc {
	return func(c core.Context) {
		req := new(listAdminMenuRequest)
		res := new(listAdminMenuResponse)
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

		searchOneData := new(admin.SearchOneData)
		searchOneData.Id = int32(ids[0])
		searchOneData.IsUsed = 1

		info, err := h.adminService.Detail(c, searchOneData)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AdminMenuListError,
				code.Text(code.AdminMenuListError)).WithError(err),
			)
			return
		}

		res.UserName = info.Username

		searchData := new(admin.SearchListMenuData)
		searchData.AdminId = int32(ids[0])

		listData, err := h.adminService.ListMenu(c, searchData)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AdminMenuListError,
				code.Text(code.AdminMenuListError)).WithError(err),
			)
			return
		}

		res.List = listData
		c.Payload(res)
	}
}
