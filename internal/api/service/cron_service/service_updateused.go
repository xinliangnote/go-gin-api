package cron_service

import (
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo"
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/cron_task_repo"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"

	"github.com/spf13/cast"
)

func (s *service) UpdateUsed(ctx core.Context, id int32, used int32) (err error) {
	data := map[string]interface{}{
		"is_used":      used,
		"updated_user": ctx.UserName(),
	}

	qb := cron_task_repo.NewQueryBuilder()
	qb.WhereId(db_repo.EqualPredicate, id)
	err = qb.Updates(s.db.GetDbW().WithContext(ctx.RequestContext()), data)
	if err != nil {
		return err
	}

	// region 操作定时任务 避免主从同步延迟，在这需要查询主库
	if used == cron_task_repo.IsUsedNo {
		s.cronServer.RemoveTask(cast.ToInt(id))
	} else {
		qb = cron_task_repo.NewQueryBuilder()
		qb.WhereId(db_repo.EqualPredicate, id)
		info, err := qb.QueryOne(s.db.GetDbW().WithContext(ctx.RequestContext()))
		if err != nil {
			return err
		}

		s.cronServer.RemoveTask(cast.ToInt(id))
		s.cronServer.AddTask(info)

	}
	// endregion

	return
}
