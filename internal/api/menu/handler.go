package menu

import (
	"github.com/xinliangnote/go-gin-api/configs"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql"
	"github.com/xinliangnote/go-gin-api/internal/repository/redis"
	"github.com/xinliangnote/go-gin-api/internal/services/menu"
	"github.com/xinliangnote/go-gin-api/pkg/hash"

	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Create 创建/编辑菜单
	// @Tags API.menu
	// @Router /api/menu [post]
	Create() core.HandlerFunc

	// Detail 菜单详情
	// @Tags API.menu
	// @Router /api/menu/{id} [get]
	Detail() core.HandlerFunc

	// Delete 删除菜单
	// @Tags API.menu
	// @Router /api/menu/{id} [delete]
	Delete() core.HandlerFunc

	// UpdateUsed 更新菜单为启用/禁用
	// @Tags API.menu
	// @Router /api/menu/used [patch]
	UpdateUsed() core.HandlerFunc

	// UpdateSort 更新菜单排序
	// @Tags API.menu
	// @Router /api/menu/sort [patch]
	UpdateSort() core.HandlerFunc

	// List 菜单列表
	// @Tags API.menu
	// @Router /api/menu [get]
	List() core.HandlerFunc

	// CreateAction 创建功能权限
	// @Tags API.menu
	// @Router /api/menu_action [post]
	CreateAction() core.HandlerFunc

	// ListAction 功能权限列表
	// @Tags API.menu
	// @Router /api/menu_action [get]
	ListAction() core.HandlerFunc

	// DeleteAction 删除功能权限
	// @Tags API.menu
	// @Router /api/menu_action/{id} [delete]
	DeleteAction() core.HandlerFunc
}

type handler struct {
	logger      *zap.Logger
	cache       redis.Repo
	hashids     hash.Hash
	menuService menu.Service
}

func New(logger *zap.Logger, db mysql.Repo, cache redis.Repo) Handler {
	return &handler{
		logger:      logger,
		cache:       cache,
		hashids:     hash.New(configs.Get().HashIds.Secret, configs.Get().HashIds.Length),
		menuService: menu.New(db, cache),
	}
}

func (h *handler) i() {}
