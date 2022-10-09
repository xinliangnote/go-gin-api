package tool

import (
	"fmt"
	"github.com/xinliangnote/go-gin-api/configs"
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

type tablesRequest struct {
	DbName string `form:"db_name"` // 数据库名称
}

type tablesResponse struct {
	List []tableData `json:"list"` // 数据表列表
}

type tableData struct {
	Name    string `json:"table_name"`    // 数据表名称
	Comment string `json:"table_comment"` // 数据表备注
}

// Tables 查询 Table
// @Summary 查询 Table
// @Description 查询 Table
// @Tags API.tool
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param db_name formData string true "数据库名称"
// @Success 200 {object} tablesResponse
// @Failure 400 {object} code.Failure
// @Router /api/tool/data/tables [post]
// @Security LoginToken
func (h *handler) Tables() core.HandlerFunc {
	return func(c core.Context) {
		req := new(tablesRequest)
		res := new(tablesResponse)
		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}

		sqlTables := getTables(req.DbName)
		// TODO 后期支持查询多个数据库
		rows, err := h.db.GetDbR().Raw(sqlTables).Rows()
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.MySQLExecError,
				code.Text(code.MySQLExecError)).WithError(err),
			)
			return
		}

		defer rows.Close()

		for rows.Next() {
			var info tableData
			err = rows.Scan(&info.Name, &info.Comment)
			if err != nil {
				fmt.Printf("execute query tables action error,had ignored, detail is [%v]\n", err.Error())
				continue
			}

			res.List = append(res.List, info)
		}

		c.Payload(res)
	}
}

func getTables(dbName string) string {
	var sqlTables string
	switch configs.Get().DataBaseType.Type {
	case "Mysql":
		sqlTables = fmt.Sprintf("SELECT `table_name`,`table_comment` FROM `information_schema`.`tables` WHERE `table_schema`= '%s'", dbName)
	case "Postgresql":
		sqlTables = `select relname as tab_name, obj_description(c.oid) as table_comment from pg_class c where obj_description(c.oid) is not null `
	}
	return sqlTables

}
