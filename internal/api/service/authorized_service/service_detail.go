package authorized_service

import (
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo"
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/authorized_repo"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

func (s *service) Detail(ctx core.Context, id int32) (info *authorized_repo.Authorized, err error) {
	qb := authorized_repo.NewQueryBuilder()
	qb.WhereIsDeleted(db_repo.EqualPredicate, -1)
	qb.WhereId(db_repo.EqualPredicate, id)

	info, err = qb.First(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}

	return
}
