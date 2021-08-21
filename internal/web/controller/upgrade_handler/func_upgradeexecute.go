package upgrade_handler

import (
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/pkg/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/web/controller/install_handler/mysql_table"
	"github.com/xinliangnote/go-gin-api/pkg/errno"
	"github.com/xinliangnote/go-gin-api/pkg/errors"
)

type upgradeExecuteRequest struct {
	TableName string `form:"table_name"` // 数据表
	Op        string `form:"op"`         // 操作类型
}

func (h *handler) UpgradeExecute() core.HandlerFunc {

	upgradeTableList := map[string]map[string]string{
		"authorized": {
			"table_sql":      mysql_table.CreateAuthorizedTableSql(),
			"table_data_sql": mysql_table.CreateAuthorizedTableDataSql(),
		},
		"authorized_api": {
			"table_sql":      mysql_table.CreateAuthorizedAPITableSql(),
			"table_data_sql": mysql_table.CreateAuthorizedAPITableDataSql(),
		},
		"admin": {
			"table_sql":      mysql_table.CreateAdminTableSql(),
			"table_data_sql": mysql_table.CreateAdminTableDataSql(),
		},
		"admin_menu": {
			"table_sql":      mysql_table.CreateAdminMenuTableSql(),
			"table_data_sql": mysql_table.CreateAdminMenuTableDataSql(),
		},
		"menu": {
			"table_sql":      mysql_table.CreateMenuTableSql(),
			"table_data_sql": mysql_table.CreateMenuTableDataSql(),
		},
		"menu_action": {
			"table_sql":      mysql_table.CreateMenuActionTableSql(),
			"table_data_sql": mysql_table.CreateMenuActionTableDataSql(),
		},
		"cron_task": {
			"table_sql":      mysql_table.CreateCronTaskTableSql(),
			"table_data_sql": "",
		},
	}

	upgradeTableOp := map[string]bool{
		"table":      true,
		"table_data": true,
	}

	return func(ctx core.Context) {
		req := new(upgradeExecuteRequest)
		if err := ctx.ShouldBindForm(req); err != nil {
			ctx.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithErr(err),
			)
			return
		}

		outPutString := ""
		db := h.db.GetDbW()

		if upgradeTableList[req.TableName] == nil {
			ctx.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.MySQLExecError,
				code.Text(code.MySQLExecError)).WithErr(errors.New("数据表不存在")),
			)
			return
		}

		if !upgradeTableOp[req.Op] {
			ctx.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.MySQLExecError,
				code.Text(code.MySQLExecError)).WithErr(errors.New("非法操作")),
			)
			return
		}

		if req.Op == "table" {
			if err := db.Exec(upgradeTableList[req.TableName]["table_sql"]).Error; err != nil {
				ctx.AbortWithError(errno.NewError(
					http.StatusBadRequest,
					code.MySQLExecError,
					code.Text(code.MySQLExecError)+" "+err.Error()).WithErr(err),
				)
				return
			}

			outPutString = "初始化 MySQL 数据表：" + req.TableName + " 成功。"
		} else if req.Op == "table_data" {
			if err := db.Exec(upgradeTableList[req.TableName]["table_data_sql"]).Error; err != nil {
				ctx.AbortWithError(errno.NewError(
					http.StatusBadRequest,
					code.MySQLExecError,
					code.Text(code.MySQLExecError)+" "+err.Error()).WithErr(err),
				)
				return
			}

			outPutString = "初始化 MySQL 数据表：" + req.TableName + " 默认数据成功。"
		}

		ctx.Payload(outPutString)
	}
}
