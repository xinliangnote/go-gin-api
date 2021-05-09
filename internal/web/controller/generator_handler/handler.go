package generator_handler

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/cache"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/db"

	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	HandlerView() core.HandlerFunc
	HandlerExecute() core.HandlerFunc

	GormView() core.HandlerFunc
	GormExecute() core.HandlerFunc
}

type handler struct {
	db     db.Repo
	logger *zap.Logger
	cache  cache.Repo
}

func New(logger *zap.Logger, db db.Repo, cache cache.Repo) Handler {
	return &handler{
		logger: logger,
		cache:  cache,
		db:     db,
	}
}

func (h *handler) i() {}
