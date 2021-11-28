package authorized

import (
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/services/authorized"
)

type createRequest struct {
	BusinessKey       string `form:"business_key"`       // 调用方key
	BusinessDeveloper string `form:"business_developer"` // 调用方对接人
	Remark            string `form:"remark"`             // 备注
}

type createResponse struct {
	Id int32 `json:"id"` // 主键ID
}

// Create 新增调用方
// @Summary 新增调用方
// @Description 新增调用方
// @Tags API.authorized
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param business_key formData string true "调用方key"
// @Param business_developer formData string true "调用方对接人"
// @Param remark formData string true "备注"
// @Success 200 {object} createResponse
// @Failure 400 {object} code.Failure
// @Router /api/authorized [post]
// @Security LoginToken
func (h *handler) Create() core.HandlerFunc {
	return func(c core.Context) {
		req := new(createRequest)
		res := new(createResponse)
		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}

		createData := new(authorized.CreateAuthorizedData)
		createData.BusinessKey = req.BusinessKey
		createData.BusinessDeveloper = req.BusinessDeveloper
		createData.Remark = req.Remark

		id, err := h.authorizedService.Create(c, createData)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AuthorizedCreateError,
				code.Text(code.AuthorizedCreateError)).WithError(err),
			)
			return
		}

		res.Id = id
		c.Payload(res)
	}
}
