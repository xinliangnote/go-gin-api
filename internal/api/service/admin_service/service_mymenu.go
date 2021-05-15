package admin_service

import (
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo"
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/admin_menu_repo"
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/menu_repo"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

type SearchMyMenuData struct {
	AdminId int32 `json:"admin_id"` // 管理员ID
}

type ListMyMenuData struct {
	Id   int32  `json:"id"`   // ID
	Pid  int32  `json:"pid"`  // 父类ID
	Name string `json:"name"` // 菜单名称
	Link string `json:"link"` // 链接地址
	Icon string `json:"icon"` // 图标
}

func (s *service) MyMenu(ctx core.Context, searchData *SearchMyMenuData) (menuData []ListMyMenuData, err error) {
	adminMenuQb := admin_menu_repo.NewQueryBuilder()
	if searchData.AdminId != 0 {
		adminMenuQb.WhereAdminId(db_repo.EqualPredicate, searchData.AdminId)
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

	menuQb := menu_repo.NewQueryBuilder()
	menuQb.WhereIsDeleted(db_repo.EqualPredicate, -1)
	menuListData, err := menuQb.
		OrderBySort(true).
		QueryAll(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}

	if len(menuListData) <= 0 {
		return
	}

	for _, menuAllV := range menuListData {
		for _, v := range adminMenuListData {
			if menuAllV.Id == v.MenuId {
				data := ListMyMenuData{
					Id:   menuAllV.Id,
					Pid:  menuAllV.Pid,
					Name: menuAllV.Name,
					Link: menuAllV.Link,
					Icon: menuAllV.Icon,
				}

				menuData = append(menuData, data)
			}
		}
	}

	return
}
