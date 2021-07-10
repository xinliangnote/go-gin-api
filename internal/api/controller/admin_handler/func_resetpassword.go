package admin_handler

import (
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/pkg/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/pkg/errno"
)

type resetPasswordRequest struct {
	Id string `uri:"id"` // HashID
}

type resetPasswordResponse struct {
	Id int32 `json:"id"` // 主键ID
}

// ResetPassword 重置密码
// @Summary 重置密码
// @Description 重置密码
// @Tags API.admin
// @Accept json
// @Produce json
// @Param id path string true "hashId"
// @Success 200 {object} resetPasswordResponse
// @Failure 400 {object} code.Failure
// @Router /api/admin/reset_password/{id} [patch]
func (h *handler) ResetPassword() core.HandlerFunc {
	return func(c core.Context) {
		req := new(resetPasswordRequest)
		res := new(resetPasswordResponse)
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

		err = h.adminService.ResetPassword(c, id)
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.AdminResetPasswordError,
				code.Text(code.AdminResetPasswordError)).WithErr(err),
			)
			return
		}

		res.Id = id
		c.Payload(res)
	}
}
