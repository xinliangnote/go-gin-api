package user_handler

import (
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/api/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/pkg/errno"
)

type deleteRequest struct {
	Id int32 `uri:"id"` // 用户ID
}

type deleteResponse struct {
	Id int32 `json:"id"` // 用户主键ID
}

// 删除用户
// @Summary 删除用户
// @Description 删除用户
// @Tags User
// @Accept  json
// @Produce  json
// @Param id path int true "用户ID"
// @Param Authorization header string true "签名"
// @Success 200 {object} deleteResponse
// @Failure 400 {object} code.Failure
// @Failure 401 {object} code.Failure
// @Router /user/delete/{id} [patch]
func (h *handler) Delete() core.HandlerFunc {
	return func(c core.Context) {
		req := new(deleteRequest)
		res := new(deleteResponse)
		if err := c.ShouldBindURI(req); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithErr(err),
			)
			return
		}

		err := h.userService.Delete(c, req.Id)
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
