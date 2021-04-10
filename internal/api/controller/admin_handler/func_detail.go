package admin_handler

import (
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/api/code"
	"github.com/xinliangnote/go-gin-api/internal/api/service/admin_service"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/pkg/errno"

	"github.com/spf13/cast"
)

type detailResponse struct {
	Username string `json:"username"` // 用户名
	Nickname string `json:"nickname"` // 昵称
	Mobile   string `json:"mobile"`   // 手机号
}

// Detail 管理员详情
// @Summary 管理员详情
// @Description 管理员详情
// @Tags API.admin
// @Accept json
// @Produce json
// @Success 200 {object} detailResponse
// @Failure 400 {object} code.Failure
// @Router /api/admin/info [get]
func (h *handler) Detail() core.HandlerFunc {
	return func(c core.Context) {
		res := new(detailResponse)

		searchOneData := new(admin_service.SearchOneData)
		searchOneData.Id = cast.ToInt32(c.UserID())
		searchOneData.IsUsed = 1

		info, err := h.adminService.Detail(c, searchOneData)
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.AdminLoginError,
				code.Text(code.AdminLoginError)).WithErr(err),
			)
			return
		}

		res.Username = info.Username
		res.Nickname = info.Nickname
		res.Mobile = info.Mobile
		c.Payload(res)
	}
}
