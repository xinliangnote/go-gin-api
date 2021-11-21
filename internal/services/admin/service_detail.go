package admin

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql/admin"
)

type SearchOneData struct {
	Id       int32  // 用户ID
	Username string // 用户名
	Nickname string // 昵称
	Mobile   string // 手机号
	Password string // 密码
	IsUsed   int32  // 是否启用 1:是  -1:否
}

func (s *service) Detail(ctx core.Context, searchOneData *SearchOneData) (info *admin.Admin, err error) {

	qb := admin.NewQueryBuilder()
	qb.WhereIsDeleted(mysql.EqualPredicate, -1)

	if searchOneData.Id != 0 {
		qb.WhereId(mysql.EqualPredicate, searchOneData.Id)
	}

	if searchOneData.Username != "" {
		qb.WhereUsername(mysql.EqualPredicate, searchOneData.Username)
	}

	if searchOneData.Nickname != "" {
		qb.WhereNickname(mysql.EqualPredicate, searchOneData.Nickname)
	}

	if searchOneData.Mobile != "" {
		qb.WhereMobile(mysql.EqualPredicate, searchOneData.Mobile)
	}

	if searchOneData.Password != "" {
		qb.WherePassword(mysql.EqualPredicate, searchOneData.Password)
	}

	if searchOneData.IsUsed != 0 {
		qb.WhereIsUsed(mysql.EqualPredicate, searchOneData.IsUsed)
	}

	info, err = qb.QueryOne(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}

	return
}
