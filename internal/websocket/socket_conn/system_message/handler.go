package system_message

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/cache"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/db"
	"github.com/xinliangnote/go-gin-api/internal/websocket/socket_server"
	"github.com/xinliangnote/go-gin-api/pkg/errors"

	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

var server socket_server.Server

type Handler interface {
	i()

	// Connect 建立 Socket 连接
	Connect() core.HandlerFunc
}

type handler struct {
	logger *zap.Logger
	cache  cache.Repo
	db     db.Repo
}

func New(logger *zap.Logger, db db.Repo, cache cache.Repo) Handler {
	return &handler{
		logger: logger,
		cache:  cache,
		db:     db,
	}
}

func GetConn() (socket_server.Server, error) {
	if server != nil {
		return server, nil
	}

	return nil, errors.New("conn is nil")
}

func (h *handler) i() {}
