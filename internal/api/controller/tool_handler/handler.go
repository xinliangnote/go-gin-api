package tool_handler

import (
	"github.com/xinliangnote/go-gin-api/configs"
	"github.com/xinliangnote/go-gin-api/internal/pkg/cache"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/db"
	"github.com/xinliangnote/go-gin-api/pkg/hash"

	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// HashIdsEncode HashIds 加密
	// @Tags API.tool
	// @Router /api/tool/hashids/encode/{id} [get]
	HashIdsEncode() core.HandlerFunc

	// HashIdsDecode HashIds 解密
	// @Tags API.tool
	// @Router /api/tool/hashids/decode/{id} [get]
	HashIdsDecode() core.HandlerFunc
}

type handler struct {
	logger  *zap.Logger
	cache   cache.Repo
	hashids hash.Hash
}

func New(logger *zap.Logger, db db.Repo, cache cache.Repo) Handler {
	return &handler{
		logger:  logger,
		cache:   cache,
		hashids: hash.New(configs.Get().HashIds.Secret, configs.Get().HashIds.Length),
	}
}

func (h *handler) i() {}
