package cron

import (
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/validation"
	"github.com/xinliangnote/go-gin-api/pkg/errno"
)

type updateUsedRequest struct {
	Id   string `form:"id"`   // 主键ID
	Used int32  `form:"used"` // 是否启用 1:是 -1:否
}

type updateUsedResponse struct {
	Id int32 `json:"id"` // 主键ID
}

// UpdateUsed 更新任务为启用/禁用
// @Summary 更新任务为启用/禁用
// @Description 更新任务为启用/禁用
// @Tags API.cron
// @Accept multipart/form-data
// @Produce json
// @Param id formData string true "Hashid"
// @Param used formData int true "是否启用 1:是 -1:否"
// @Success 200 {object} updateUsedResponse
// @Failure 400 {object} code.Failure
// @Router /api/cron/used [patch]
func (h *handler) UpdateUsed() core.HandlerFunc {
	return func(ctx core.Context) {
		req := new(updateUsedRequest)
		res := new(updateUsedResponse)
		if err := ctx.ShouldBindForm(req); err != nil {
			ctx.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ParamBindError,
				validation.Error(err)).WithErr(err),
			)
			return
		}

		ids, err := h.hashids.HashidsDecode(req.Id)
		if err != nil {
			ctx.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.HashIdsDecodeError,
				code.Text(code.HashIdsDecodeError)).WithErr(err),
			)
			return
		}

		id := int32(ids[0])

		err = h.cronService.UpdateUsed(ctx, id, req.Used)
		if err != nil {
			ctx.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.AdminUpdateError,
				code.Text(code.AdminUpdateError)).WithErr(err),
			)
			return
		}

		res.Id = id
		ctx.Payload(res)
	}
}
