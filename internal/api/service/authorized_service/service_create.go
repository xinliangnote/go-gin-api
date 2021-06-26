package authorized_service

import (
	"crypto/rand"
	"encoding/hex"
	"io"

	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/authorized_repo"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

type CreateAuthorizedData struct {
	BusinessKey       string `json:"business_key"`       // 调用方key
	BusinessDeveloper string `json:"business_developer"` // 调用方对接人
	Remark            string `json:"remark"`             // 备注
}

func (s *service) Create(ctx core.Context, authorizedData *CreateAuthorizedData) (id int32, err error) {
	buf := make([]byte, 10)
	io.ReadFull(rand.Reader, buf)
	secret := hex.EncodeToString(buf)

	model := authorized_repo.NewModel()
	model.BusinessKey = authorizedData.BusinessKey
	model.BusinessSecret = secret
	model.BusinessDeveloper = authorizedData.BusinessDeveloper
	model.Remark = authorizedData.Remark
	model.CreatedUser = ctx.UserName()
	model.IsUsed = 1
	model.IsDeleted = -1

	id, err = model.Create(s.db.GetDbW().WithContext(ctx.RequestContext()))
	if err != nil {
		return 0, err
	}
	return
}
