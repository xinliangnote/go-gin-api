package cron

import "github.com/spf13/cast"

func (s *server) RemoveTask(taskId int) {
	name := cast.ToString(taskId)
	s.cron.RemoveJob(name)
}
