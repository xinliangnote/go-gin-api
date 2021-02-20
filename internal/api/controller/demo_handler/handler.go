package demo_handler

import (
	"github.com/xinliangnote/go-gin-api/internal/api/service/user_service"
	"github.com/xinliangnote/go-gin-api/internal/pkg/cache"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/db"
	"github.com/xinliangnote/go-gin-api/internal/pkg/grpc"

	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	// i 为了避免被其他包实现
	i()
	// 示例：支持 get 请求的方法
	Get() core.HandlerFunc
	// 示例：支持 post 请求的方法
	Post() core.HandlerFunc
	// 获取授权信息
	Auth() core.HandlerFunc
	// Trace 示例
	Trace() core.HandlerFunc
}

type handler struct {
	logger      *zap.Logger
	cache       cache.Repo
	grpConn     grpc.ClientConn
	userService user_service.UserService
}

func New(logger *zap.Logger, db db.Repo, cache cache.Repo, grpConn grpc.ClientConn) Handler {
	return &handler{
		logger:      logger,
		cache:       cache,
		grpConn:     grpConn,
		userService: user_service.NewUserService(db, cache),
	}
}

func (h *handler) i() {}
