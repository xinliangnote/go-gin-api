package cron_service

import (
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo"
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/cron_task_repo"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

type SearchData struct {
	Page     int    // 第几页
	PageSize int    // 每页显示条数
	Name     string // 任务名称
	Protocol int32  // 执行方式
	IsUsed   int32  // 是否启用
}

func (s *service) PageList(ctx core.Context, searchData *SearchData) (listData []*cron_task_repo.CronTask, err error) {
	page := searchData.Page
	if page == 0 {
		page = 1
	}

	pageSize := searchData.PageSize
	if pageSize == 0 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	qb := cron_task_repo.NewQueryBuilder()

	if searchData.Name != "" {
		qb.WhereName(db_repo.EqualPredicate, searchData.Name)
	}

	if searchData.Protocol != 0 {
		qb.WhereProtocol(db_repo.EqualPredicate, searchData.Protocol)
	}

	if searchData.IsUsed != 0 {
		qb.WhereIsUsed(db_repo.EqualPredicate, searchData.IsUsed)
	}

	listData, err = qb.
		Limit(pageSize).
		Offset(offset).
		OrderById(false).
		QueryAll(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}

	return
}
