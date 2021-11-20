package upgrade

import (
	"github.com/xinliangnote/go-gin-api/internal/api/repository/redis"
	"github.com/xinliangnote/go-gin-api/internal/pkg/db"

	"go.uber.org/zap"
)

type handler struct {
	db     db.Repo
	logger *zap.Logger
	cache  redis.Repo
}

func New(logger *zap.Logger, db db.Repo, cache redis.Repo) *handler {
	return &handler{
		logger: logger,
		cache:  cache,
		db:     db,
	}
}
