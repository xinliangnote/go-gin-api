package user_handler

import (
	"github.com/xinliangnote/go-gin-api/internal/api/service/user_service"
	"github.com/xinliangnote/go-gin-api/internal/pkg/cache"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/db"

	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	// i 为了避免被其他包实现
	i()
	// Create 创建用户
	Create() core.HandlerFunc
	// UpdateNickNameByID 编辑用户 - 通过主键ID更新用户昵称
	UpdateNickNameByID() core.HandlerFunc
	// Delete 删除用户 - 通过主键ID更新 is_deleted = 1
	Delete() core.HandlerFunc
	// Detail 用户详情
	Detail() core.HandlerFunc
}

type handler struct {
	logger      *zap.Logger
	cache       cache.Repo
	userService user_service.UserService
}

func New(logger *zap.Logger, db db.Repo, cache cache.Repo) Handler {
	return &handler{
		logger:      logger,
		cache:       cache,
		userService: user_service.NewUserService(db, cache),
	}
}

func (h *handler) i() {}
