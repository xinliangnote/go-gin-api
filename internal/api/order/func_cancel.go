package order

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

type cancelRequest struct{}

type cancelResponse struct{}

// Cancel 取消订单
// @Summary 取消订单
// @Description 取消订单
// @Tags API.order
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body cancelRequest true "请求信息"
// @Success 200 {object} cancelResponse
// @Failure 400 {object} code.Failure
// @Router /api/order/cancel [post]
func (h *handler) Cancel() core.HandlerFunc {
	return func(ctx core.Context) {

	}
}
