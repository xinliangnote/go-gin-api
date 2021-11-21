package socket

import "go.uber.org/zap"

func (s *server) OnMessage() {
	defer func() {
		s.OnClose()
	}()

	for {
		//接收消息
		_, message, err := s.socket.ReadMessage()
		if err != nil {
			s.logger.Error("socket on message error", zap.Error(err))
			break
		}

		// 为了便于演示，仅输出到日志文件
		s.logger.Info("receive message: " + string(message))
	}
}
