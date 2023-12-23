package generator_handler

import (
	"github.com/xinliangnote/go-gin-api/internal/repository/iface"
	"github.com/xinliangnote/go-gin-api/internal/repository/redis"

	"go.uber.org/zap"
)

type handler struct {
	db     iface.Repo
	logger *zap.Logger
	cache  redis.Repo
}

func New(logger *zap.Logger, db iface.Repo, cache redis.Repo) *handler {
	return &handler{
		logger: logger,
		cache:  cache,
		db:     db,
	}
}
