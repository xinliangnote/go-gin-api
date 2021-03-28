package authorized_service

import (
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo"
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/authorized_repo"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

func (s *service) PageListCount(ctx core.Context, searchData *SearchData) (total int64, err error) {
	qb := authorized_repo.NewQueryBuilder()
	qb = qb.WhereIsDeleted(db_repo.EqualPredicate, -1)

	if searchData.BusinessKey != "" {
		qb.WhereBusinessKey(db_repo.EqualPredicate, searchData.BusinessKey)
	}

	if searchData.BusinessSecret != "" {
		qb.WhereBusinessSecret(db_repo.EqualPredicate, searchData.BusinessSecret)
	}

	if searchData.BusinessDeveloper != "" {
		qb.WhereBusinessDeveloper(db_repo.EqualPredicate, searchData.BusinessDeveloper)
	}

	total, err = qb.Count(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return 0, err
	}

	return
}
