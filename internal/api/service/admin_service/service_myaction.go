package admin_service

import (
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo"
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/admin_menu_repo"
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/menu_action_repo"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

type SearchMyActionData struct {
	AdminId int32 `json:"admin_id"` // 管理员ID
}

func (s *service) MyAction(ctx core.Context, searchData *SearchMyActionData) (actionData []*menu_action_repo.MenuAction, err error) {
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

	var menuIds []int32
	for _, v := range adminMenuListData {
		menuIds = append(menuIds, v.MenuId)
	}

	actionQb := menu_action_repo.NewQueryBuilder()
	actionQb.WhereIsDeleted(db_repo.EqualPredicate, -1)
	actionQb.WhereMenuIdIn(menuIds)
	actionData, err = actionQb.QueryAll(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}

	return
}
