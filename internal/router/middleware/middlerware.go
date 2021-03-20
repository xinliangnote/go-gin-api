package middleware

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/cache"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/pkg/errno"

	"go.uber.org/zap"
)

var _ Middleware = (*middleware)(nil)

type Middleware interface {
	// i 为了避免被其他包实现
	i()

	// JWT 中间件
	Jwt(ctx core.Context) (userId int64, userName string, err errno.Error)

	// Resubmit 中间件
	Resubmit() core.HandlerFunc

	// DisableLog 不记录日志
	DisableLog() core.HandlerFunc
}

type middleware struct {
	logger *zap.Logger
	cache  cache.Repo
}

func New(logger *zap.Logger, cache cache.Repo) Middleware {
	return &middleware{
		logger: logger,
		cache:  cache,
	}
}

func (m *middleware) i() {}
