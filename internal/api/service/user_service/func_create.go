package user_service

import (
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/user_demo_repo"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

type CreateUserInfo struct {
	UserName string `json:"user_name"` // 用户名
	NickName string `json:"nick_name"` // 昵称
	Mobile   string `json:"mobile"`    // 手机号
}

func (u *userSer) Create(ctx core.Context, user *CreateUserInfo) (id int32, err error) {
	model := user_demo_repo.NewModel()
	model.UserName = user.UserName
	model.NickName = user.NickName
	model.Mobile = user.Mobile

	id, err = model.Create(u.db.GetDbW().WithContext(ctx.RequestContext()))
	if err != nil {
		return 0, err
	}
	return
}
