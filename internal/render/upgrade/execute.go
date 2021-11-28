package upgrade

import (
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/proposal/tablesqls"
	"github.com/xinliangnote/go-gin-api/pkg/errors"
)

type upgradeExecuteRequest struct {
	TableName string `form:"table_name"` // 数据表
	Op        string `form:"op"`         // 操作类型
}

func (h *handler) UpgradeExecute() core.HandlerFunc {

	upgradeTableList := map[string]map[string]string{
		"authorized": {
			"table_sql":      tablesqls.CreateAuthorizedTableSql(),
			"table_data_sql": tablesqls.CreateAuthorizedTableDataSql(),
		},
		"authorized_api": {
			"table_sql":      tablesqls.CreateAuthorizedAPITableSql(),
			"table_data_sql": tablesqls.CreateAuthorizedAPITableDataSql(),
		},
		"admin": {
			"table_sql":      tablesqls.CreateAdminTableSql(),
			"table_data_sql": tablesqls.CreateAdminTableDataSql(),
		},
		"admin_menu": {
			"table_sql":      tablesqls.CreateAdminMenuTableSql(),
			"table_data_sql": tablesqls.CreateAdminMenuTableDataSql(),
		},
		"menu": {
			"table_sql":      tablesqls.CreateMenuTableSql(),
			"table_data_sql": tablesqls.CreateMenuTableDataSql(),
		},
		"menu_action": {
			"table_sql":      tablesqls.CreateMenuActionTableSql(),
			"table_data_sql": tablesqls.CreateMenuActionTableDataSql(),
		},
		"cron_task": {
			"table_sql":      tablesqls.CreateCronTaskTableSql(),
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
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}

		outPutString := ""
		db := h.db.GetDbW()

		if upgradeTableList[req.TableName] == nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.MySQLExecError,
				code.Text(code.MySQLExecError)).WithError(errors.New("数据表不存在")),
			)
			return
		}

		if !upgradeTableOp[req.Op] {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.MySQLExecError,
				code.Text(code.MySQLExecError)).WithError(errors.New("非法操作")),
			)
			return
		}

		if req.Op == "table" {
			if err := db.Exec(upgradeTableList[req.TableName]["table_sql"]).Error; err != nil {
				ctx.AbortWithError(core.Error(
					http.StatusBadRequest,
					code.MySQLExecError,
					code.Text(code.MySQLExecError)+" "+err.Error()).WithError(err),
				)
				return
			}

			outPutString = "初始化 MySQL 数据表：" + req.TableName + " 成功。"
		} else if req.Op == "table_data" {
			if err := db.Exec(upgradeTableList[req.TableName]["table_data_sql"]).Error; err != nil {
				ctx.AbortWithError(core.Error(
					http.StatusBadRequest,
					code.MySQLExecError,
					code.Text(code.MySQLExecError)+" "+err.Error()).WithError(err),
				)
				return
			}

			outPutString = "初始化 MySQL 数据表：" + req.TableName + " 默认数据成功。"
		}

		ctx.Payload(outPutString)
	}
}
