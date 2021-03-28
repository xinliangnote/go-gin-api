package authorized_service

import (
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/authorized_repo"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

func (s *service) Delete(ctx core.Context, id int32) (err error) {
	model := authorized_repo.NewModel()
	model.Id = id

	data := map[string]interface{}{
		"is_deleted":   1,
		"updated_user": "system", // TODO
	}

	err = model.Updates(s.db.GetDbW().WithContext(ctx.RequestContext()), data)
	if err != nil {
		return err
	}

	return
}
