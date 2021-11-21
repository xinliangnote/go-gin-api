package socket

import (
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

func (s *server) OnSend(message []byte) error {
	err := s.socket.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		s.OnClose()
		s.logger.Error("socket on send error", zap.Error(err))
	}
	return err
}
