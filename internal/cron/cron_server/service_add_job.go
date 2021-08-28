package cron_server

import (
	"fmt"

	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/cron_task_repo"

	"github.com/jakecoffman/cron"
)

func (s *server) AddJob(task *cron_task_repo.CronTask) cron.FuncJob {
	return func() {
		s.taskCount.Add()
		defer s.taskCount.Done()

		msg := fmt.Sprintf("开始执行任务：(%d)%s [%s]", task.Id, task.Name, task.Spec)
		s.logger.Info(msg)
	}
}
