package admin

import (
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/services/admin"
	"github.com/xinliangnote/go-gin-api/pkg/errno"

	"github.com/spf13/cast"
)

type modifyPersonalInfoRequest struct {
	Nickname string `form:"nickname"` // 昵称
	Mobile   string `form:"mobile"`   // 手机号
}

type modifyPersonalInfoResponse struct {
	Username string `json:"username"` // 用户账号
}

// ModifyPersonalInfo 修改个人信息
// @Summary 修改个人信息
// @Description 修改个人信息
// @Tags API.admin
// @Accept multipart/form-data
// @Produce json
// @Param nickname formData string true "昵称"
// @Param mobile formData string true "手机号"
// @Success 200 {object} modifyPersonalInfoResponse
// @Failure 400 {object} code.Failure
// @Router /api/admin/modify_personal_info [patch]
func (h *handler) ModifyPersonalInfo() core.HandlerFunc {
	return func(c core.Context) {
		req := new(modifyPersonalInfoRequest)
		res := new(modifyPersonalInfoResponse)
		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithErr(err),
			)
			return
		}

		userId := cast.ToInt32(c.UserID())

		modifyData := new(admin.ModifyData)
		modifyData.Nickname = req.Nickname
		modifyData.Mobile = req.Mobile

		if err := h.adminService.ModifyPersonalInfo(c, userId, modifyData); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.AdminModifyPersonalInfoError,
				code.Text(code.AdminModifyPersonalInfoError)).WithErr(err),
			)
			return
		}

		res.Username = c.UserName()
		c.Payload(res)
	}
}
