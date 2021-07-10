package admin_handler

import (
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/api/service/admin_service"
	"github.com/xinliangnote/go-gin-api/internal/pkg/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/password"
	"github.com/xinliangnote/go-gin-api/pkg/errno"

	"github.com/spf13/cast"
)

type modifyPasswordRequest struct {
	OldPassword string `form:"old_password"` // 旧密码
	NewPassword string `form:"new_password"` // 新密码
}

type modifyPasswordResponse struct {
	Username string `json:"username"` // 用户账号
}

// ModifyPassword 修改密码
// @Summary 修改密码
// @Description 修改密码
// @Tags API.admin
// @Accept multipart/form-data
// @Produce json
// @Param old_password formData string true "旧密码"
// @Param new_password formData string true "新密码"
// @Success 200 {object} modifyPasswordResponse
// @Failure 400 {object} code.Failure
// @Router /api/admin/modify_password [patch]
func (h *handler) ModifyPassword() core.HandlerFunc {
	return func(c core.Context) {
		req := new(modifyPasswordRequest)
		res := new(modifyPasswordResponse)
		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithErr(err),
			)
			return
		}

		userId := cast.ToInt32(c.UserID())

		searchOneData := new(admin_service.SearchOneData)
		searchOneData.Id = userId
		searchOneData.Password = password.GeneratePassword(req.OldPassword)
		searchOneData.IsUsed = 1

		info, err := h.adminService.Detail(c, searchOneData)
		if err != nil || info == nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.AdminModifyPasswordError,
				code.Text(code.AdminModifyPasswordError)).WithErr(err),
			)
			return
		}

		if err := h.adminService.ModifyPassword(c, userId, req.NewPassword); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.AdminModifyPasswordError,
				code.Text(code.AdminModifyPasswordError)).WithErr(err),
			)
			return
		}

		res.Username = c.UserName()
		c.Payload(res)
	}
}
