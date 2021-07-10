package tool_handler

import (
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/pkg/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/pkg/errno"
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
// @Accept json
// @Produce json
// @Param id path string true "需解密的密文"
// @Success 200 {object} hashIdsDecodeResponse
// @Failure 400 {object} code.Failure
// @Router /api/tool/hashids/decode/{id} [get]
func (h *handler) HashIdsDecode() core.HandlerFunc {
	return func(c core.Context) {
		req := new(hashIdsDecodeRequest)
		res := new(hashIdsDecodeResponse)
		if err := c.ShouldBindURI(req); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithErr(err),
			)
			return
		}

		hashId, err := h.hashids.HashidsDecode(req.Id)
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.HashIdsDecodeError,
				code.Text(code.HashIdsDecodeError)).WithErr(err),
			)
			return
		}

		res.Val = hashId[0]

		c.Payload(res)
	}
}
