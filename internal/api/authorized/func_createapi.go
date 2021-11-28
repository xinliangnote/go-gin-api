package authorized

import (
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/services/authorized"
)

type createAPIRequest struct {
	Id     string `form:"id"`     // HashID
	Method string `form:"method"` // 请求方法
	API    string `form:"api"`    // 请求地址
}

type createAPIResponse struct {
	Id int32 `json:"id"` // 主键ID
}

// CreateAPI 授权调用方接口地址
// @Summary 授权调用方接口地址
// @Description 授权调用方接口地址
// @Tags API.authorized
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param id formData string true "HashID"
// @Param method formData string true "请求方法"
// @Param api formData string true "请求地址"
// @Success 200 {object} createAPIResponse
// @Failure 400 {object} code.Failure
// @Router /api/authorized_api [post]
// @Security LoginToken
func (h *handler) CreateAPI() core.HandlerFunc {
	return func(c core.Context) {
		req := new(createAPIRequest)
		res := new(createAPIResponse)
		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}

		ids, err := h.hashids.HashidsDecode(req.Id)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.HashIdsDecodeError,
				code.Text(code.HashIdsDecodeError)).WithError(err),
			)
			return
		}

		id := int32(ids[0])

		// 通过 id 查询出 business_key
		authorizedInfo, err := h.authorizedService.Detail(c, id)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AuthorizedDetailError,
				code.Text(code.AuthorizedDetailError)).WithError(err),
			)
			return
		}

		createAPIData := new(authorized.CreateAuthorizedAPIData)
		createAPIData.BusinessKey = authorizedInfo.BusinessKey
		createAPIData.Method = req.Method
		createAPIData.API = req.API

		createId, err := h.authorizedService.CreateAPI(c, createAPIData)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AuthorizedCreateAPIError,
				code.Text(code.AuthorizedCreateAPIError)).WithError(err),
			)
			return
		}

		res.Id = createId
		c.Payload(res)
	}
}
