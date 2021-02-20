package user_handler

import (
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/api/code"
	"github.com/xinliangnote/go-gin-api/internal/api/service/user_service"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/pkg/errno"

	"github.com/pkg/errors"
)

type createRequest struct {
	UserName string `json:"user_name"` // 用户名
	NickName string `json:"nick_name"` // 昵称
	Mobile   string `json:"mobile"`    // 手机号
}

type createResponse struct {
	Id uint `json:"id"` // 主键ID
}

// 创建用户
// @Summary 创建用户
// @Description 创建用户
// @Tags User
// @Accept  json
// @Produce  json
// @Param Request body createRequest true "请求信息"
// @Param Authorization header string true "签名"
// @Success 200 {object} createResponse
// @Failure 400 {object} code.Failure
// @Failure 401 {object} code.Failure
// @Router /user/create [post]
func (h *handler) Create() core.HandlerFunc {
	return func(c core.Context) {
		req := new(createRequest)
		res := new(createResponse)
		if err := c.ShouldBindJSON(req); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithErr(err),
			)
			return
		}

		if req.UserName == "" {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.IllegalUserName,
				code.Text(code.IllegalUserName)).WithErr(errors.New("req.UserName = ''")),
			)
			return
		}

		createUserData := new(user_service.CreateUserInfo)
		createUserData.Mobile = req.Mobile
		createUserData.NickName = req.NickName
		createUserData.UserName = req.UserName

		id, err := h.userService.Create(c, createUserData)
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.UserCreateError,
				code.Text(code.UserCreateError)).WithErr(err),
			)
			return
		}

		res.Id = id
		c.Payload(res)

	}
}
