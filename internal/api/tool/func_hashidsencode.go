package tool

import (
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"

	"github.com/spf13/cast"
)

type hashIdsEncodeRequest struct {
	Id int32 `uri:"id"` // 需加密的数字
}

type hashIdsEncodeResponse struct {
	Val string `json:"val"` // 加密后的值
}

// HashIdsEncode HashIds 加密
// @Summary HashIds 加密
// @Description HashIds 加密
// @Tags API.tool
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param id path string true "需加密的数字"
// @Success 200 {object} hashIdsEncodeResponse
// @Failure 400 {object} code.Failure
// @Router /api/tool/hashids/encode/{id} [get]
// @Security LoginToken
func (h *handler) HashIdsEncode() core.HandlerFunc {
	return func(c core.Context) {
		req := new(hashIdsEncodeRequest)
		res := new(hashIdsEncodeResponse)
		if err := c.ShouldBindURI(req); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}

		hashId, err := h.hashids.HashidsEncode([]int{cast.ToInt(req.Id)})
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.HashIdsEncodeError,
				code.Text(code.HashIdsEncodeError)).WithError(err),
			)
			return
		}

		res.Val = hashId

		c.Payload(res)
	}
}
