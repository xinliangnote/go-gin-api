package router

import (
	"github.com/xinliangnote/go-gin-api/internal/websocket/sysmessage"
)

func setSocketRouter(r *resource) {
	systemMessage := sysmessage.New(r.logger, r.db, r.cache)

	// 无需记录日志
	socket := r.mux.Group("/socket", r.middles.DisableLog())
	{
		// 系统消息
		socket.GET("/system/message", systemMessage.Connect())
	}
}
