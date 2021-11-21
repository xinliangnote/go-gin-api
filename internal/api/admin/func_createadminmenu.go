package admin

import (
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/services/admin"
	"github.com/xinliangnote/go-gin-api/pkg/errno"
)

type createAdminMenuRequest struct {
	Id      string `form:"id"`      // HashID
	Actions string `form:"actions"` // 功能权限ID,多个用,分割
}

type createAdminMenuResponse struct {
	Id int32 `json:"id"` // 主键ID
}

// CreateAdminMenu 提交菜单授权
// @Summary 提交菜单授权
// @Description 提交菜单授权
// @Tags API.admin
// @Accept multipart/form-data
// @Produce json
// @Param id formData string true "Hashid"
// @Param actions formData string true "功能权限ID,多个用,分割"
// @Success 200 {object} createResponse
// @Failure 400 {object} code.Failure
// @Router /api/admin/menu [post]
func (h *handler) CreateAdminMenu() core.HandlerFunc {
	return func(c core.Context) {
		req := new(createAdminMenuRequest)
		res := new(createAdminMenuResponse)
		if err := c.ShouldBindForm(req); err != nil {
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

		createData := new(admin.CreateMenuData)
		createData.AdminId = int32(ids[0])
		createData.Actions = req.Actions

		err = h.adminService.CreateMenu(c, createData)
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.AdminMenuCreateError,
				code.Text(code.AdminMenuCreateError)).WithErr(err),
			)
			return
		}

		res.Id = int32(ids[0])
		c.Payload(res)
	}
}
