package admin_service

import (
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo"
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/admin_repo"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

type SearchOneData struct {
	Id       int32  // 用户ID
	Username string // 用户名
	Nickname string // 昵称
	Mobile   string // 手机号
	Password string // 密码
	IsUsed   int32  // 是否启用 1:是  -1:否
}

func (s *service) Detail(ctx core.Context, searchOneData *SearchOneData) (info *admin_repo.Admin, err error) {

	qb := admin_repo.NewQueryBuilder()
	qb.WhereIsDeleted(db_repo.EqualPredicate, -1)

	if searchOneData.Id != 0 {
		qb.WhereId(db_repo.EqualPredicate, searchOneData.Id)
	}

	if searchOneData.Username != "" {
		qb.WhereUsername(db_repo.EqualPredicate, searchOneData.Username)
	}

	if searchOneData.Nickname != "" {
		qb.WhereNickname(db_repo.EqualPredicate, searchOneData.Nickname)
	}

	if searchOneData.Mobile != "" {
		qb.WhereMobile(db_repo.EqualPredicate, searchOneData.Mobile)
	}

	if searchOneData.Password != "" {
		qb.WherePassword(db_repo.EqualPredicate, searchOneData.Password)
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
