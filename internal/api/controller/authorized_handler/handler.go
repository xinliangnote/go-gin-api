package authorized_handler

import (
	"github.com/xinliangnote/go-gin-api/configs"
	"github.com/xinliangnote/go-gin-api/internal/api/service/authorized_service"
	"github.com/xinliangnote/go-gin-api/internal/pkg/cache"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/db"
	"github.com/xinliangnote/go-gin-api/pkg/hash"

	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Create 新增调用方
	// @Tags API.authorized
	// @Router /api/authorized [post]
	Create() core.HandlerFunc

	// CreateAPI 授权调用方接口地址
	// @Tags API.authorized
	// @Router /api/authorized_api [post]
	CreateAPI() core.HandlerFunc

	// List 调用方列表
	// @Tags API.authorized
	// @Router /api/authorized [get]
	List() core.HandlerFunc

	// ListAPI 调用方接口地址列表
	// @Tags API.authorized
	// @Router /api/authorized_api [get]
	ListAPI() core.HandlerFunc

	// Delete 删除调用方
	// @Tags API.authorized
	// @Router /api/authorized/{id} [delete]
	Delete() core.HandlerFunc

	// DeleteAPI 删除调用方接口地址
	// @Tags API.authorized
	// @Router /api/authorized_api/{id} [delete]
	DeleteAPI() core.HandlerFunc

	// UpdateUsed 更新调用方为启用/禁用
	// @Tags API.authorized
	// @Router /api/authorized/used [patch]
	UpdateUsed() core.HandlerFunc
}

type handler struct {
	logger            *zap.Logger
	cache             cache.Repo
	authorizedService authorized_service.Service
	hashids           hash.Hash
}

func New(logger *zap.Logger, db db.Repo, cache cache.Repo) Handler {
	return &handler{
		logger:            logger,
		cache:             cache,
		authorizedService: authorized_service.New(db, cache),
		hashids:           hash.New(configs.Get().HashIds.Secret, configs.Get().HashIds.Length),
	}
}

func (h *handler) i() {}
