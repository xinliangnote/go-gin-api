package admin

import (
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/password"
	"github.com/xinliangnote/go-gin-api/internal/services/admin"
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
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param old_password formData string true "旧密码"
// @Param new_password formData string true "新密码"
// @Success 200 {object} modifyPasswordResponse
// @Failure 400 {object} code.Failure
// @Router /api/admin/modify_password [patch]
// @Security LoginToken
func (h *handler) ModifyPassword() core.HandlerFunc {
	return func(ctx core.Context) {
		req := new(modifyPasswordRequest)
		res := new(modifyPasswordResponse)
		if err := ctx.ShouldBindForm(req); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}

		searchOneData := new(admin.SearchOneData)
		searchOneData.Id = ctx.SessionUserInfo().UserID
		searchOneData.Password = password.GeneratePassword(req.OldPassword)
		searchOneData.IsUsed = 1

		info, err := h.adminService.Detail(ctx, searchOneData)
		if err != nil || info == nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AdminModifyPasswordError,
				code.Text(code.AdminModifyPasswordError)).WithError(err),
			)
			return
		}

		if err := h.adminService.ModifyPassword(ctx, ctx.SessionUserInfo().UserID, req.NewPassword); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AdminModifyPasswordError,
				code.Text(code.AdminModifyPasswordError)).WithError(err),
			)
			return
		}

		res.Username = ctx.SessionUserInfo().UserName
		ctx.Payload(res)
	}
}
