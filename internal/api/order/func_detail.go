package order

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

type detailRequest struct{}

type detailResponse struct{}

// Detail 取消订单
// @Summary 取消订单
// @Description 取消订单
// @Tags API.order
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body detailRequest true "请求信息"
// @Success 200 {object} detailResponse
// @Failure 400 {object} code.Failure
// @Router /api/order/{id} [get]
func (h *handler) Detail() core.HandlerFunc {
	return func(ctx core.Context) {

	}
}
