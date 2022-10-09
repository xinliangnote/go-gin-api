package order

import (
	"github.com/xinliangnote/go-gin-api/configs"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/repository/iface"
	"github.com/xinliangnote/go-gin-api/internal/repository/redis"
	"github.com/xinliangnote/go-gin-api/pkg/hash"

	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Create 创建订单
	// @Tags API.order
	// @Router /api/order/create [post]
	Create() core.HandlerFunc

	// Cancel 取消订单
	// @Tags API.order
	// @Router /api/order/cancel [post]
	Cancel() core.HandlerFunc

	// Detail 取消订单
	// @Tags API.order
	// @Router /api/order/{id} [get]
	Detail() core.HandlerFunc
}

type handler struct {
	logger  *zap.Logger
	db      iface.Repo
	cache   redis.Repo
	hashids hash.Hash
}

func New(logger *zap.Logger, db iface.Repo, cache redis.Repo) Handler {
	return &handler{
		logger:  logger,
		db:      db,
		cache:   cache,
		hashids: hash.New(configs.Get().HashIds.Secret, configs.Get().HashIds.Length),
	}
}

func (h *handler) i() {}
