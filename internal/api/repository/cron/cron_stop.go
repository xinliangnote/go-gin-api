package cron

func (s *server) Stop() {
	s.cron.Stop()
	s.taskCount.Exit()
}
