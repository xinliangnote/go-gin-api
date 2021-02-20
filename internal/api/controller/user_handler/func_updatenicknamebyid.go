package user_handler

import (
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/api/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/pkg/errno"
)

type updateNickNameByIDRequest struct {
	Id       uint   `json:"id"`        // 用户主键ID
	NickName string `json:"nick_name"` // 昵称
}

type updateNickNameByIDResponse struct {
	Id uint `json:"id"` // 用户主键ID
}

// 编辑用户 - 通过用户主键ID更新用户昵称
// @Summary 编辑用户 - 通过用户主键ID更新用户昵称
// @Description 编辑用户 - 通过用户主键ID更新用户昵称
// @Tags User
// @Accept  json
// @Produce  json
// @Param Request body updateNickNameByIDRequest true "请求信息"
// @Param Authorization header string true "签名"
// @Success 200 {object} updateNickNameByIDResponse
// @Failure 400 {object} code.Failure
// @Failure 401 {object} code.Failure
// @Router /user/update [put]
func (h *handler) UpdateNickNameByID() core.HandlerFunc {
	return func(c core.Context) {
		req := new(updateNickNameByIDRequest)
		res := new(updateNickNameByIDResponse)
		if err := c.ShouldBindJSON(req); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithErr(err),
			)
			return
		}

		err := h.userService.UpdateNickNameByID(c, req.Id, req.NickName)
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.UserUpdateError,
				code.Text(code.UserUpdateError)).WithErr(err),
			)
			return
		}

		res.Id = req.Id

		c.Payload(res)
	}
}
