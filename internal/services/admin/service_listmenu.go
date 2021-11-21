package admin

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql/admin_menu"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql/menu"
)

type SearchListMenuData struct {
	AdminId int32 `json:"admin_id"` // 管理员ID
}

type ListMenuData struct {
	Id     int32  `json:"id"`      // ID
	Pid    int32  `json:"pid"`     // 父类ID
	Name   string `json:"name"`    // 菜单名称
	IsHave int32  `json:"is_have"` // 是否已拥有权限
}

func (s *service) ListMenu(ctx core.Context, searchData *SearchListMenuData) (menuData []ListMenuData, err error) {
	menuQb := menu.NewQueryBuilder()
	menuQb.WhereIsDeleted(mysql.EqualPredicate, -1)
	menuListData, err := menuQb.
		OrderBySort(true).
		QueryAll(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}

	if len(menuListData) <= 0 {
		return
	}

	menuData = make([]ListMenuData, len(menuListData))
	for k, v := range menuListData {
		data := ListMenuData{
			Id:     v.Id,
			Pid:    v.Pid,
			Name:   v.Name,
			IsHave: 0,
		}

		menuData[k] = data
	}

	adminMenuQb := admin_menu.NewQueryBuilder()
	if searchData.AdminId != 0 {
		adminMenuQb.WhereAdminId(mysql.EqualPredicate, searchData.AdminId)
	}

	adminMenuListData, err := adminMenuQb.
		OrderById(false).
		QueryAll(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}

	if len(adminMenuListData) <= 0 {
		return
	}

	for k, v := range menuData {
		for _, haveV := range adminMenuListData {
			if haveV.MenuId == v.Id {
				menuData[k].IsHave = 1
			}
		}
	}

	return
}
