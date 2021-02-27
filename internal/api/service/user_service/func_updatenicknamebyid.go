package user_service

import (
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/user_demo_repo"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

func (u *userSer) UpdateNickNameByID(ctx core.Context, id int32, nickname string) (err error) {
	model := user_demo_repo.NewModel()
	model.Id = id

	data := map[string]interface{}{
		"nick_name": nickname,
	}

	err = model.Updates(u.db.GetDbW().WithContext(ctx.RequestContext()), data)
	if err != nil {
		return err
	}

	return nil
}
