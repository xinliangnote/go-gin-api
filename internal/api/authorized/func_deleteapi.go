package authorized

import (
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/pkg/errno"
)

type deleteAPIRequest struct {
	Id string `uri:"id"` // HashID
}

type deleteAPIResponse struct {
	Id int32 `json:"id"` // 主键ID
}

// DeleteAPI 删除调用方接口地址
// @Summary 删除调用方接口地址
// @Description 删除调用方接口地址
// @Tags API.authorized
// @Accept json
// @Produce json
// @Param id path string true "主键ID"
// @Success 200 {object} deleteAPIResponse
// @Failure 400 {object} code.Failure
// @Router /api/authorized_api/{id} [delete]
func (h *handler) DeleteAPI() core.HandlerFunc {
	return func(c core.Context) {
		req := new(deleteAPIRequest)
		res := new(deleteAPIResponse)
		if err := c.ShouldBindURI(req); err != nil {
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

		err = h.authorizedService.DeleteAPI(c, id)
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.AuthorizedDeleteAPIError,
				code.Text(code.AuthorizedDeleteAPIError)).WithErr(err),
			)
			return
		}

		res.Id = id
		c.Payload(res)
	}
}
