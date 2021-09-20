package socket_server

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/cache"
	"github.com/xinliangnote/go-gin-api/internal/pkg/db"
	"github.com/xinliangnote/go-gin-api/pkg/errors"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

var _ Server = (*server)(nil)

type server struct {
	logger *zap.Logger
	db     db.Repo
	cache  cache.Repo
	socket *websocket.Conn
}

type Server interface {
	i()

	// OnMessage 接收消息
	OnMessage()

	// OnSend 发送消息
	OnSend(message []byte) error

	// OnClose 关闭
	OnClose()
}

func New(logger *zap.Logger, db db.Repo, cache cache.Repo, conn *websocket.Conn) (Server, error) {
	if logger == nil {
		return nil, errors.New("logger required")
	}

	if db == nil {
		return nil, errors.New("db required")
	}

	if cache == nil {
		return nil, errors.New("cache required")
	}

	if conn == nil {
		return nil, errors.New("conn required")
	}

	return &server{
		logger: logger,
		db:     db,
		cache:  cache,
		socket: conn,
	}, nil
}

func (s *server) i() {}
