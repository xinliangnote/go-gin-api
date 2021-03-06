package user_service

import (
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/user_demo_repo"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

func (u *userSer) Delete(ctx core.Context, id int32) (err error) {
	model := user_demo_repo.NewModel()
	model.Id = id
	err = model.Delete(u.db.GetDbW().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil
	}
	return nil
}
