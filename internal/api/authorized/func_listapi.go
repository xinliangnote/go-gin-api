package authorized

import (
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/services/authorized"

	"github.com/spf13/cast"
)

type listAPIRequest struct {
	Id string `form:"id"` // hashID
}

type listAPIData struct {
	HashId      string `json:"hash_id"`      // hashID
	BusinessKey string `json:"business_key"` // 调用方key
	Method      string `json:"method"`       // 调用方secret
	API         string `json:"api"`          // 调用方对接人
}

type listAPIResponse struct {
	BusinessKey string        `json:"business_key"` // 调用方key
	List        []listAPIData `json:"list"`
}

// ListAPI 调用方接口地址列表
// @Summary 调用方接口地址列表
// @Description 调用方接口地址列表
// @Tags API.authorized
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param id query string true "hashID"
// @Success 200 {object} listAPIResponse
// @Failure 400 {object} code.Failure
// @Router /api/authorized_api [get]
// @Security LoginToken
func (h *handler) ListAPI() core.HandlerFunc {
	return func(c core.Context) {
		req := new(listAPIRequest)
		res := new(listAPIResponse)
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

		res.BusinessKey = authorizedInfo.BusinessKey

		searchAPIData := new(authorized.SearchAPIData)
		searchAPIData.BusinessKey = authorizedInfo.BusinessKey

		resListData, err := h.authorizedService.ListAPI(c, searchAPIData)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AuthorizedListAPIError,
				code.Text(code.AuthorizedListAPIError)).WithError(err),
			)
			return
		}

		res.List = make([]listAPIData, len(resListData))

		for k, v := range resListData {
			hashId, err := h.hashids.HashidsEncode([]int{cast.ToInt(v.Id)})
			if err != nil {
				c.AbortWithError(core.Error(
					http.StatusBadRequest,
					code.HashIdsEncodeError,
					code.Text(code.HashIdsEncodeError)).WithError(err),
				)
				return
			}

			data := listAPIData{
				HashId:      hashId,
				BusinessKey: v.BusinessKey,
				Method:      v.Method,
				API:         v.Api,
			}

			res.List[k] = data
		}

		c.Payload(res)
	}
}
