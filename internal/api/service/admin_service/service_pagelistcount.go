package admin_service

import (
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo"
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/admin_repo"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

func (s *service) PageListCount(ctx core.Context, searchData *SearchData) (total int64, err error) {
	qb := admin_repo.NewQueryBuilder()
	qb = qb.WhereIsDeleted(db_repo.EqualPredicate, -1)

	if searchData.Username != "" {
		qb.WhereUsername(db_repo.EqualPredicate, searchData.Username)
	}

	if searchData.Nickname != "" {
		qb.WhereNickname(db_repo.EqualPredicate, searchData.Nickname)
	}

	if searchData.Mobile != "" {
		qb.WhereMobile(db_repo.EqualPredicate, searchData.Mobile)
	}

	total, err = qb.Count(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return 0, err
	}

	return
}
