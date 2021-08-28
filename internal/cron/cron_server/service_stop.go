package cron_server

func (s *server) Stop() {
	s.cron.Stop()
	s.taskCount.Exit()
}
