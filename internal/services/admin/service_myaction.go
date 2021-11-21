package admin

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql/admin_menu"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql/menu_action"
)

type SearchMyActionData struct {
	AdminId int32 `json:"admin_id"` // 管理员ID
}

type MyActionData struct {
	Id     int32  // 主键
	MenuId int32  // 菜单栏ID
	Method string // 请求方式
	Api    string // 请求地址
}

func (s *service) MyAction(ctx core.Context, searchData *SearchMyActionData) (actionData []MyActionData, err error) {
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

	var menuIds []int32
	for _, v := range adminMenuListData {
		menuIds = append(menuIds, v.MenuId)
	}

	actionQb := menu_action.NewQueryBuilder()
	actionQb.WhereIsDeleted(mysql.EqualPredicate, -1)
	actionQb.WhereMenuIdIn(menuIds)
	actionListData, err := actionQb.QueryAll(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}

	if len(actionListData) <= 0 {
		return
	}

	actionData = make([]MyActionData, len(actionListData))

	for k, v := range actionListData {
		data := MyActionData{
			Id:     v.Id,
			MenuId: v.MenuId,
			Method: v.Method,
			Api:    v.Api,
		}

		actionData[k] = data
	}

	return
}
