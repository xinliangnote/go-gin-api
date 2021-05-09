package menu_service

import (
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo"
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/menu_repo"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

type UpdateMenuData struct {
	Name string // 菜单名称
	Link string // 链接地址
	Icon string // 图标
}

func (s *service) Modify(ctx core.Context, id int32, menuData *UpdateMenuData) (err error) {
	data := map[string]interface{}{
		"name":         menuData.Name,
		"link":         menuData.Link,
		"icon":         menuData.Icon,
		"updated_user": ctx.UserName(),
	}

	qb := menu_repo.NewQueryBuilder()
	qb.WhereId(db_repo.EqualPredicate, id)
	err = qb.Updates(s.db.GetDbW().WithContext(ctx.RequestContext()), data)
	if err != nil {
		return err
	}

	return
}
