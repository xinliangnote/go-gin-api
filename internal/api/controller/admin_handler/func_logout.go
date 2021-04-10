package admin_handler

import (
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/api/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/cache"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/password"
	"github.com/xinliangnote/go-gin-api/pkg/errno"

	"github.com/pkg/errors"
	"github.com/spf13/cast"
)

type logoutResponse struct {
	Username string `json:"username"` // 用户账号
}

// Logout 管理员登出
// @Summary 管理员登出
// @Description 管理员登出
// @Tags API.admin
// @Accept json
// @Produce json
// @Success 200 {object} logoutResponse
// @Failure 400 {object} code.Failure
// @Router /api/admin/login [post]
func (h *handler) Logout() core.HandlerFunc {
	return func(c core.Context) {
		res := new(logoutResponse)
		res.Username = c.UserName()

		if !h.cache.Del(h.adminService.CacheKeyPrefix()+password.GenerateLoginToken(cast.ToInt32(c.UserID())), cache.WithTrace(c.Trace())) {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.AdminLogOutError,
				code.Text(code.AdminLogOutError)).WithErr(errors.New("cache del err")),
			)
			return
		}

		c.Payload(res)
	}
}
