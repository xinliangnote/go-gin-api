package cron

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql/cron_task"
)

type SearchOneData struct {
	Id int32 // 任务ID
}

func (s *service) Detail(ctx core.Context, searchOneData *SearchOneData) (info *cron_task.CronTask, err error) {
	qb := cron_task.NewQueryBuilder()

	if searchOneData.Id != 0 {
		qb.WhereId(mysql.EqualPredicate, searchOneData.Id)
	}

	info, err = qb.QueryOne(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}

	return
}
