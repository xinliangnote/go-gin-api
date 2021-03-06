package demo_handler

import (
	"net/http"
	"time"

	"github.com/xinliangnote/go-gin-api/configs"
	"github.com/xinliangnote/go-gin-api/internal/api/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/pkg/errno"
	"github.com/xinliangnote/go-gin-api/pkg/token"
)

type authResponse struct {
	Authorization string `json:"authorization"` // 签名
	ExpireTime    int64  `json:"expire_time"`   // 过期时间
}

// 获取授权信息
// @Summary 获取授权信息
// @Description 获取授权信息
// @Tags Demo
// @Accept  json
// @Produce  json
// @Success 200 {object} authResponse
// @Failure 400 {object} code.Failure
// @Router /auth/get [post]
func (h *handler) Auth() core.HandlerFunc {
	return func(c core.Context) {
		cfg := configs.Get().JWT
		tokenString, err := token.New(cfg.Secret).JwtSign(1, "xinliangnote", time.Hour*cfg.ExpireDuration)
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.AuthorizationError,
				code.Text(code.AuthorizationError)).WithErr(err),
			)
			return
		}

		res := new(authResponse)
		res.Authorization = tokenString
		res.ExpireTime = time.Now().Add(time.Hour * cfg.ExpireDuration).Unix()

		c.Payload(res)
	}
}
