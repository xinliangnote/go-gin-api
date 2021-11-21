package socket

import "go.uber.org/zap"

func (s *server) OnClose() {
	err := s.socket.Close()
	if err != nil {
		s.logger.Error("socket on closed error", zap.Error(err))
	}
}
