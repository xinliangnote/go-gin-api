package user_handler

import (
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/api/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/pkg/ddm"
	"github.com/xinliangnote/go-gin-api/pkg/errno"
)

type detailRequest struct {
	UserName string `uri:"username"` // 用户名
}

type detailResponse struct {
	Id       uint       `json:"id"`        // 用户主键ID
	UserName string     `json:"user_name"` // 用户名
	NickName string     `json:"nick_name"` // 昵称
	Mobile   ddm.Mobile `json:"mobile"`    // 手机号（脱敏）
}

// 用户详情
// @Summary 用户详情
// @Description 用户详情
// @Tags User
// @Accept  json
// @Produce  json
// @Param username path string true "用户名"
// @Param Authorization header string true "签名"
// @Success 200 {object} detailResponse
// @Failure 400 {object} code.Failure
// @Failure 401 {object} code.Failure
// @Router /user/info/{username} [get]
func (h *handler) Detail() core.HandlerFunc {
	return func(c core.Context) {
		req := new(detailRequest)
		res := new(detailResponse)
		if err := c.ShouldBindURI(req); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithErr(err),
			)
			return
		}

		user, err := h.userService.GetUserByUserName(c, req.UserName)
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.UserSearchError,
				code.Text(code.UserSearchError)).WithErr(err),
			)
			return
		}

		res.Id = user.Id
		res.UserName = user.UserName
		res.NickName = user.NickName
		res.Mobile = ddm.Mobile(user.Mobile)

		c.Payload(res)
	}
}
