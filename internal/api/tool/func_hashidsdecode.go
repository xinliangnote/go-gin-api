package tool

import (
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

type hashIdsDecodeRequest struct {
	Id string `uri:"id"` // 需解密的密文
}

type hashIdsDecodeResponse struct {
	Val int `json:"val"` // 解密后的值
}

// HashIdsDecode HashIds 解密
// @Summary HashIds 解密
// @Description HashIds 解密
// @Tags API.tool
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param id path string true "需解密的密文"
// @Success 200 {object} hashIdsDecodeResponse
// @Failure 400 {object} code.Failure
// @Router /api/tool/hashids/decode/{id} [get]
// @Security LoginToken
func (h *handler) HashIdsDecode() core.HandlerFunc {
	return func(c core.Context) {
		req := new(hashIdsDecodeRequest)
		res := new(hashIdsDecodeResponse)
		if err := c.ShouldBindURI(req); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}

		hashId, err := h.hashids.HashidsDecode(req.Id)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.HashIdsDecodeError,
				code.Text(code.HashIdsDecodeError)).WithError(err),
			)
			return
		}

		res.Val = hashId[0]

		c.Payload(res)
	}
}
