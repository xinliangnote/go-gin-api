package menu_service

import (
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo"
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/menu_repo"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

type SearchOneData struct {
	Id     int32 // 用户ID
	IsUsed int32 // 是否启用 1:是  -1:否
}

func (s *service) Detail(ctx core.Context, searchOneData *SearchOneData) (info *menu_repo.Menu, err error) {

	qb := menu_repo.NewQueryBuilder()
	qb.WhereIsDeleted(db_repo.EqualPredicate, -1)

	if searchOneData.Id != 0 {
		qb.WhereId(db_repo.EqualPredicate, searchOneData.Id)
	}

	if searchOneData.IsUsed != 0 {
		qb.WhereIsUsed(db_repo.EqualPredicate, searchOneData.IsUsed)
	}

	info, err = qb.QueryOne(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}

	return
}
