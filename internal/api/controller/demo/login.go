package demo

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/errno"
	"github.com/xinliangnote/go-gin-api/internal/pkg/token"

	"go.uber.org/zap"
)

type loginRequest struct {
	UserID   int    `json:"user_id" form:"user_id"`     // 用户ID（>0）
	UserName string `json:"user_name" form:"user_name"` // 用户名
}

type loginResponse struct {
	Authorization string `json:"authorization"` // 签名
}

// 登录获取 Authorization 码
// @Summary 登录获取 Authorization 码
// @Description 登录获取 Authorization 码
// @Tags Demo
// @Accept  json
// @Produce  json
// @Param loginRequest body loginRequest true "请求信息"
// @Success 200 {object} loginResponse "签名信息"
// @Router /user/login [post]
func (d *Demo) Login() core.HandlerFunc {
	return func(c core.Context) {
		req := new(loginRequest)
		res := new(loginResponse)
		if err := c.ShouldBindJSON(req); err != nil {
			c.SetPayload(errno.ErrParam)
			return
		}

		tokenString, err := token.Sign(req.UserID, req.UserName)
		if err != nil {
			d.logger.Error("token sign err", zap.Error(err))
			res.Authorization = ""
		} else {
			res.Authorization = tokenString
		}
		c.SetPayload(errno.OK.WithData(res))
	}
}
