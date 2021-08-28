package cron_service

import (
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo"
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/cron_task_repo"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

func (s *service) Execute(ctx core.Context, id int32) (err error) {
	qb := cron_task_repo.NewQueryBuilder()
	qb.WhereId(db_repo.EqualPredicate, id)
	info, err := qb.QueryOne(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return err
	}

	info.Spec = "手动执行"
	go s.cronServer.AddJob(info)()

	return nil
}
