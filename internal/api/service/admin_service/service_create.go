package admin_service

import (
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/admin_repo"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/password"
)

type CreateAdminData struct {
	Username string // 用户名
	Nickname string // 昵称
	Mobile   string // 手机号
	Password string // 密码
}

func (s *service) Create(ctx core.Context, adminData *CreateAdminData) (id int32, err error) {
	model := admin_repo.NewModel()
	model.Username = adminData.Username
	model.Password = password.GeneratePassword(adminData.Password)
	model.Nickname = adminData.Nickname
	model.Mobile = adminData.Mobile
	model.CreatedUser = ctx.UserName()
	model.IsUsed = 1
	model.IsDeleted = -1

	id, err = model.Create(s.db.GetDbW().WithContext(ctx.RequestContext()))
	if err != nil {
		return 0, err
	}
	return
}
