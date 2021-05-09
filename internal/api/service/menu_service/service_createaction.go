package menu_service

import (
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/menu_action_repo"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

type CreateMenuActionData struct {
	MenuId int32  `json:"menu_id"` // 菜单栏ID
	Method string `json:"method"`  // 请求方法
	API    string `json:"api"`     // 请求地址
}

func (s *service) CreateAction(ctx core.Context, menuActionData *CreateMenuActionData) (id int32, err error) {
	model := menu_action_repo.NewModel()
	model.MenuId = menuActionData.MenuId
	model.Method = menuActionData.Method
	model.Api = menuActionData.API
	model.CreatedUser = ctx.UserName()
	model.IsDeleted = -1

	id, err = model.Create(s.db.GetDbW().WithContext(ctx.RequestContext()))
	if err != nil {
		return 0, err
	}
	return
}
