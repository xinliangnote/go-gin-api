package menu_service

import (
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/menu_repo"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

func (s *service) UpdateUsed(ctx core.Context, id int32, used int32) (err error) {
	model := menu_repo.NewModel()
	model.Id = id

	data := map[string]interface{}{
		"is_used":      used,
		"updated_user": ctx.UserName(),
	}

	err = model.Updates(s.db.GetDbW().WithContext(ctx.RequestContext()), data)
	if err != nil {
		return err
	}

	return
}
