package tool

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/xinliangnote/go-gin-api/internal/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"

	"github.com/spf13/cast"
)

type tableColumn struct {
	ColumnName    string `json:"column_name"`    // 字段名称
	ColumnComment string `json:"column_comment"` // 字段注释
}

type searchMySQLRequest struct {
	DbName    string `form:"db_name"`    // 数据库名称
	TableName string `form:"table_name"` // 数据表名称
	SQL       string `form:"sql"`        // SQL 语句
}

type searchMySQLResponse struct {
	Cols     []string                 `json:"cols"`      // 查询后的行
	ColsInfo []tableColumn            `json:"cols_info"` // 行信息
	List     []map[string]interface{} `json:"list"`      // 查询后的数据
}

var preFilterList = map[string]bool{
	"insert": true,
	"update": true,
	"delete": true,
	"create": true,
	"source": true,
	"rename": true,
}

var whiteListKeyword = []string{
	"is_deleted",
	"updated_at",
	"created_at",
	"updated_user",
	"created_user",
	"show create table",
}

var filterListKeyword = []string{
	"insert",
	"update",
	"truncate",
	"delete",
	"create",
	"alter",
	"rename",
	"drop",
	"replace",
	"sleep",
	"grant",
	"revoke",
	"load_file",
	"outfile",
	"transaction",
	"commit",
	"mysqldump",
	"into",
}

// SearchMySQL 执行 SQL 语句
// @Summary 执行 SQL 语句
// @Description 执行 SQL 语句
// @Tags API.tool
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param db_name formData string true "数据库名称"
// @Param table_name formData string true "数据表名称"
// @Param sql formData string true "SQL 语句"
// @Success 200 {object} searchMySQLResponse
// @Failure 400 {object} code.Failure
// @Router /api/tool/data/mysql [post]
// @Security LoginToken
func (h *handler) SearchMySQL() core.HandlerFunc {
	return func(c core.Context) {
		req := new(searchMySQLRequest)
		res := new(searchMySQLResponse)
		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}

		sql := strings.ToLower(strings.TrimSpace(req.SQL))
		if sql == "" {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.MySQLExecError,
				"SQL 语句不能为空！"),
			)
			return
		}

		if preFilterList[string([]byte(sql)[:6])] {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.MySQLExecError,
				"SQL 语句不能以 "+string([]byte(sql)[:6])+" 开头！"),
			)
			return
		}

		for _, f := range filterListKeyword {
			if find := strings.Contains(sql, f); find {

				isWhiteList := false
				for _, w := range whiteListKeyword {
					if whiteFind := strings.Contains(sql, w); whiteFind {
						isWhiteList = true
						break
					}
				}

				if !isWhiteList {
					c.AbortWithError(core.Error(
						http.StatusBadRequest,
						code.MySQLExecError,
						"SQL 语句存在敏感词： "+f+"！"),
					)
					return
				}

			}
		}

		if strings.ToLower(string([]byte(sql)[:6])) == "select" {
			sql += " LIMIT 100"
		}

		// TODO 后期支持查询多个数据库
		rows, err := h.db.GetDbR().Raw(sql).Rows()
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.MySQLExecError,
				"MySQL "+err.Error()).WithError(err),
			)
			return
		}

		defer rows.Close()

		cols, _ := rows.Columns()

		var data []map[string]interface{}

		for rows.Next() {
			// Create a slice of interface{}'s to represent each column,
			// and a second slice to contain pointers to each item in the columns slice.
			columns := make([]interface{}, len(cols))
			columnPointers := make([]interface{}, len(cols))
			for i := range columns {
				columnPointers[i] = &columns[i]
			}

			// Scan the result into the column pointers...
			if err := rows.Scan(columnPointers...); err != nil {
				fmt.Printf("query table scan error, detail is [%v]\n", err.Error())
				continue
			}

			// Create our map, and retrieve the value for each column from the pointers slice,
			// storing it in the map with the name of the column as the key.
			m := make(map[string]interface{})
			for i, colName := range cols {
				val := columnPointers[i].(*interface{})
				m[colName] = cast.ToString(*val)
			}

			data = append(data, m)

		}

		res.List = data
		res.Cols = cols

		sqlTableColumn := fmt.Sprintf("SELECT `COLUMN_NAME`, `COLUMN_COMMENT` FROM `information_schema`.`columns` WHERE `table_schema`= '%s' AND `table_name`= '%s' ORDER BY `ORDINAL_POSITION` ASC",
			req.DbName, req.TableName)

		rows, err = h.db.GetDbR().Raw(sqlTableColumn).Rows()
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.MySQLExecError,
				"MySQL "+err.Error()).WithError(err),
			)
			return
		}
		defer rows.Close()

		var tableColumns []tableColumn

		for rows.Next() {
			var column tableColumn
			err = rows.Scan(
				&column.ColumnName,
				&column.ColumnComment)

			if err != nil {
				fmt.Printf("query table column scan error, detail is [%v]\n", err.Error())
				continue
			}

			tableColumns = append(tableColumns, column)
		}

		res.ColsInfo = tableColumns

		c.Payload(res)
	}
}
