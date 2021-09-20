package system_message

import (
	"net/http"
	"time"

	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/websocket/socket_server"

	"github.com/gorilla/websocket"
)

func (h *handler) Connect() core.HandlerFunc {
	var upGrader = websocket.Upgrader{
		HandshakeTimeout: 5 * time.Second,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	return func(ctx core.Context) {
		ws, err := upGrader.Upgrade(ctx.ResponseWriter(), ctx.Request(), nil)
		if err != nil {
			return
		}

		server, err = socket_server.New(h.logger, h.db, h.cache, ws)
		if err != nil {
			return
		}

		go server.OnMessage()
	}
}
