package tool

import (
	"github.com/xinliangnote/go-gin-api/configs"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

type dbsResponse struct {
	List []dbData `json:"list"` // 数据库列表
}

type dbData struct {
	DbName string `json:"db_name"` // 数据库名称
}

// Dbs 查询 DB
// @Summary 查询 DB
// @Description 查询 DB
// @Tags API.tool
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Success 200 {object} dbsResponse
// @Failure 400 {object} code.Failure
// @Router /api/tool/data/dbs [get]
// @Security LoginToken
func (h *handler) Dbs() core.HandlerFunc {
	return func(c core.Context) {
		res := new(dbsResponse)

		// TODO 后期支持查询多个数据库
		data := dbData{
			DbName: configs.Get().MySQL.Read.Name,
		}

		res.List = append(res.List, data)
		c.Payload(res)
	}
}
