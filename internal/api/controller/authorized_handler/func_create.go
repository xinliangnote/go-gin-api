package authorized_handler

import (
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/api/code"
	"github.com/xinliangnote/go-gin-api/internal/api/service/authorized_service"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/pkg/errno"
)

type createRequest struct {
	BusinessKey       string `json:"business_key"`       // 调用方key
	BusinessDeveloper string `json:"business_developer"` // 调用方对接人
	Remark            string `json:"remark"`             // 备注
}

type createResponse struct {
	Id int32 `json:"id"` // 主键ID
}

// Create 新增调用方
// @Summary 新增调用方
// @Description 新增调用方
// @Tags API.authorized
// @Accept json
// @Produce json
// @Param Request body createRequest true "请求信息"
// @Success 200 {object} createResponse
// @Failure 400 {object} code.Failure
// @Router /api/authorized [post]
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

		createData := new(authorized_service.CreateAuthorizedData)
		createData.BusinessKey = req.BusinessKey
		createData.BusinessDeveloper = req.BusinessDeveloper
		createData.Remark = req.Remark

		id, err := h.authorizedService.Create(c, createData)
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.AuthorizedCreateError,
				code.Text(code.AuthorizedCreateError)).WithErr(err),
			)
			return
		}

		res.Id = id
		c.Payload(res)
	}
}
