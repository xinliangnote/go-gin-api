package authorized_handler

import (
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/api/code"
	"github.com/xinliangnote/go-gin-api/internal/api/service/authorized_service"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/pkg/errno"
)

type createAPIRequest struct {
	Id     string `json:"id"`     // HashID
	Method string `json:"method"` // 请求方法
	API    string `json:"api"`    // 请求地址
}

type createAPIResponse struct {
	Id int32 `json:"id"` // 主键ID
}

// CreateAPI 授权调用方接口地址
// @Summary 授权调用方接口地址
// @Description 授权调用方接口地址
// @Tags API.authorized
// @Accept json
// @Produce json
// @Param Request body createAPIRequest true "请求信息"
// @Success 200 {object} createAPIResponse
// @Failure 400 {object} code.Failure
// @Router /api/authorized_api [post]
func (h *handler) CreateAPI() core.HandlerFunc {
	return func(c core.Context) {
		req := new(createAPIRequest)
		res := new(createAPIResponse)
		if err := c.ShouldBindJSON(req); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithErr(err),
			)
			return
		}

		ids, err := h.hashids.HashidsDecode(req.Id)
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.HashIdsDecodeError,
				code.Text(code.HashIdsDecodeError)).WithErr(err),
			)
			return
		}

		id := int32(ids[0])

		// 通过 id 查询出 business_key
		authorizedInfo, err := h.authorizedService.Detail(c, id)
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.AuthorizedDetailError,
				code.Text(code.AuthorizedDetailError)).WithErr(err),
			)
			return
		}

		createAPIData := new(authorized_service.CreateAuthorizedAPIData)
		createAPIData.BusinessKey = authorizedInfo.BusinessKey
		createAPIData.Method = req.Method
		createAPIData.API = req.API

		createId, err := h.authorizedService.CreateAPI(c, createAPIData)
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.AuthorizedCreateAPIError,
				code.Text(code.AuthorizedCreateAPIError)).WithErr(err),
			)
			return
		}

		res.Id = createId
		c.Payload(res)
	}
}
