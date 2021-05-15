package admin_handler

import (
	"net/http"

	"github.com/xinliangnote/go-gin-api/configs"
	"github.com/xinliangnote/go-gin-api/internal/api/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/cache"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/password"
	"github.com/xinliangnote/go-gin-api/pkg/errno"
)

type offlineRequest struct {
	Id string `form:"id"` // 主键ID
}

type offlineResponse struct {
	Id int32 `json:"id"` // 主键ID
}

// Offline 下线管理员
// @Summary 下线管理员
// @Description 下线管理员
// @Tags API.admin
// @Accept multipart/form-data
// @Produce json
// @Param id formData string true "Hashid"
// @Success 200 {object} offlineResponse
// @Failure 400 {object} code.Failure
// @Router /api/admin/offline [patch]
func (h *handler) Offline() core.HandlerFunc {
	return func(c core.Context) {
		req := new(offlineRequest)
		res := new(offlineResponse)
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

		id := int32(ids[0])

		b := h.cache.Del(configs.RedisKeyPrefixLoginUser+password.GenerateLoginToken(id), cache.WithTrace(c.Trace()))
		if !b {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.AdminOfflineError,
				code.Text(code.AdminOfflineError)),
			)
			return
		}

		res.Id = id
		c.Payload(res)
	}
}
