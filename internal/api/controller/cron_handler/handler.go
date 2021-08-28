package cron_handler

import (
	"github.com/xinliangnote/go-gin-api/configs"
	"github.com/xinliangnote/go-gin-api/internal/api/service/cron_service"
	"github.com/xinliangnote/go-gin-api/internal/cron/cron_server"
	"github.com/xinliangnote/go-gin-api/internal/pkg/cache"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/db"
	"github.com/xinliangnote/go-gin-api/pkg/hash"

	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Create 创建任务
	// @Tags API.cron
	// @Router /api/cron [post]
	Create() core.HandlerFunc

	// Modify 编辑任务
	// @Tags API.cron
	// @Router /api/cron/{id} [post]
	Modify() core.HandlerFunc

	// List 任务列表
	// @Tags API.cron
	// @Router /api/cron [get]
	List() core.HandlerFunc

	// UpdateUsed 更新任务为启用/禁用
	// @Tags API.cron
	// @Router /api/cron/used [patch]
	UpdateUsed() core.HandlerFunc

	// Detail 获取单条任务详情
	// @Tags API.cron
	// @Router /api/cron/:id [get]
	Detail() core.HandlerFunc

	// Execute 手动执行任务
	// @Tags API.cron
	// @Router /api/cron/:id [patch]
	Execute() core.HandlerFunc
}

type handler struct {
	logger      *zap.Logger
	cache       cache.Repo
	hashids     hash.Hash
	cronService cron_service.Service
}

func New(logger *zap.Logger, db db.Repo, cache cache.Repo, cron cron_server.Server) Handler {
	return &handler{
		logger:      logger,
		cache:       cache,
		hashids:     hash.New(configs.Get().HashIds.Secret, configs.Get().HashIds.Length),
		cronService: cron_service.New(db, cache, cron),
	}
}

func (h *handler) i() {}
