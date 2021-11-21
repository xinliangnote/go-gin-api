package cron

import (
	"fmt"
	"math"

	"github.com/xinliangnote/go-gin-api/internal/repository/mysql"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql/cron_task"

	"go.uber.org/zap"
)

func (s *server) Start() {
	s.cron.Start()
	go s.taskCount.Wait()

	qb := cron_task.NewQueryBuilder()
	qb.WhereIsUsed(mysql.EqualPredicate, cron_task.IsUsedYES)
	totalNum, err := qb.Count(s.db.GetDbR())
	if err != nil {
		s.logger.Fatal("cron initialize tasks count err", zap.Error(err))
	}

	pageSize := 50
	maxPage := int(math.Ceil(float64(totalNum) / float64(pageSize)))

	taskNum := 0
	s.logger.Info("开始初始化后台任务")

	for page := 1; page <= maxPage; page++ {
		qb = cron_task.NewQueryBuilder()
		qb.WhereIsUsed(mysql.EqualPredicate, cron_task.IsUsedYES)
		listData, err := qb.
			Limit(pageSize).
			Offset((page - 1) * pageSize).
			OrderById(false).
			QueryAll(s.db.GetDbR())
		if err != nil {
			s.logger.Fatal("cron initialize tasks list err", zap.Error(err))
		}

		for _, item := range listData {
			s.AddTask(item)
			taskNum++
		}
	}

	s.logger.Info(fmt.Sprintf("后台任务初始化完成，总数量：%d", taskNum))
}
