package admin_service

import (
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/admin_repo"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

type ModifyData struct {
	Nickname string // 昵称
	Mobile   string // 手机号
}

func (s *service) ModifyPersonalInfo(ctx core.Context, id int32, modifyData *ModifyData) (err error) {
	model := admin_repo.NewModel()
	model.Id = id

	data := map[string]interface{}{
		"nickname":     modifyData.Nickname,
		"mobile":       modifyData.Mobile,
		"updated_user": ctx.UserName(),
	}

	err = model.Updates(s.db.GetDbW().WithContext(ctx.RequestContext()), data)
	if err != nil {
		return err
	}

	return
}
