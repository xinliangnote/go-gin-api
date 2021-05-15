package menu_service

import (
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo"
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/menu_repo"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

type SearchData struct {
	Pid int32 // 父类ID
}

func (s *service) List(ctx core.Context, searchData *SearchData) (listData []*menu_repo.Menu, err error) {

	qb := menu_repo.NewQueryBuilder()
	qb.WhereIsDeleted(db_repo.EqualPredicate, -1)

	if searchData.Pid != 0 {
		qb.WherePid(db_repo.EqualPredicate, searchData.Pid)
	}

	listData, err = qb.
		OrderBySort(true).
		QueryAll(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}

	return
}
