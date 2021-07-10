package admin_handler

import (
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/api/service/admin_service"
	"github.com/xinliangnote/go-gin-api/internal/pkg/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/pkg/errno"
)

type listAdminMenuRequest struct {
	Id string `uri:"id"` // HashID
}

type listAdminMenuResponse struct {
	List     []admin_service.ListMenuData `json:"list"`
	UserName string                       `json:"username"`
}

// ListAdminMenu 菜单授权列表
// @Summary 菜单授权列表
// @Description 菜单授权列表
// @Tags API.admin
// @Accept json
// @Produce json
// @Param id path string true "hashId"
// @Success 200 {object} listAdminMenuResponse
// @Failure 400 {object} code.Failure
// @Router /api/admin/menu/:id [get]
func (h *handler) ListAdminMenu() core.HandlerFunc {
	return func(c core.Context) {
		req := new(listAdminMenuRequest)
		res := new(listAdminMenuResponse)
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

		searchOneData := new(admin_service.SearchOneData)
		searchOneData.Id = int32(ids[0])
		searchOneData.IsUsed = 1

		info, err := h.adminService.Detail(c, searchOneData)
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.AdminMenuListError,
				code.Text(code.AdminMenuListError)).WithErr(err),
			)
			return
		}

		res.UserName = info.Username

		searchData := new(admin_service.SearchListMenuData)
		searchData.AdminId = int32(ids[0])

		listData, err := h.adminService.ListMenu(c, searchData)
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.AdminMenuListError,
				code.Text(code.AdminMenuListError)).WithErr(err),
			)
			return
		}

		res.List = listData
		c.Payload(res)
	}
}
