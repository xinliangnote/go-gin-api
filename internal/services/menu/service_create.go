package menu

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql/menu"
)

type CreateMenuData struct {
	Pid   int32  // 父类ID
	Name  string // 菜单名称
	Link  string // 链接地址
	Icon  string // 图标
	Level int32  // 菜单类型 1:一级菜单 2:二级菜单
}

func (s *service) Create(ctx core.Context, menuData *CreateMenuData) (id int32, err error) {
	model := menu.NewModel()
	model.Pid = menuData.Pid
	model.Name = menuData.Name
	model.Link = menuData.Link
	model.Icon = menuData.Icon
	model.Level = menuData.Level
	model.CreatedUser = ctx.SessionUserInfo().UserName
	model.IsUsed = 1
	model.IsDeleted = -1

	id, err = model.Create(s.db.GetDbW().WithContext(ctx.RequestContext()))
	if err != nil {
		return 0, err
	}
	return
}
