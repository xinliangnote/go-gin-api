package user_service

import (
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo"
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/user_demo_repo"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

func (u *userSer) GetUserByUserName(ctx core.Context, username string) (user *user_demo_repo.UserDemo, err error) {
	user, err = user_demo_repo.NewQueryBuilder().
		WhereUserName(db_repo.EqualPredicate, username).
		QueryOne(u.db.GetDbR().WithContext(ctx.RequestContext()))

	if err != nil {
		return user, err
	}

	if user == nil {
		user = user_demo_repo.NewModel()
	}

	return user, nil
}
