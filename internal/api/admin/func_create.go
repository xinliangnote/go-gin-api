package admin

import (
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/validation"
	"github.com/xinliangnote/go-gin-api/internal/services/admin"
	"github.com/xinliangnote/go-gin-api/pkg/errno"
)

type createRequest struct {
	Username string `form:"username" binding:"required"` // 用户名
	Nickname string `form:"nickname" binding:"required"` // 昵称
	Mobile   string `form:"mobile" binding:"required"`   // 手机号
	Password string `form:"password" binding:"required"` // 密码
}

type createResponse struct {
	Id int32 `json:"id"` // 主键ID
}

// Create 新增管理员
// @Summary 新增管理员
// @Description 新增管理员
// @Tags API.admin
// @Accept multipart/form-data
// @Produce json
// @Param username formData string true "用户名"
// @Param nickname formData string true "昵称"
// @Param mobile formData string true "手机号"
// @Param password formData string true "密码"
// @Success 200 {object} createResponse
// @Failure 400 {object} code.Failure
// @Router /api/admin [post]
func (h *handler) Create() core.HandlerFunc {
	return func(c core.Context) {
		req := new(createRequest)
		res := new(createResponse)
		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ParamBindError,
				validation.Error(err)).WithErr(err),
			)
			return
		}

		createData := new(admin.CreateAdminData)
		createData.Nickname = req.Nickname
		createData.Username = req.Username
		createData.Mobile = req.Mobile
		createData.Password = req.Password

		id, err := h.adminService.Create(c, createData)
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.AdminCreateError,
				code.Text(code.AdminCreateError)).WithErr(err),
			)
			return
		}

		res.Id = id
		c.Payload(res)
	}
}
