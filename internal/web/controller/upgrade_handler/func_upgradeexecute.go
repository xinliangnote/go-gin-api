package upgrade_handler

import (
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/api/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/web/controller/install_handler/mysql_table"
	"github.com/xinliangnote/go-gin-api/pkg/errno"
)

type upgradeExecuteRequest struct {
	TableName string `form:"table_name"` // 数据表
	Op        string `form:"op"`         // 操作类型
}

func (h *handler) UpgradeExecute() core.HandlerFunc {
	return func(c core.Context) {
		req := new(upgradeExecuteRequest)
		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithErr(err),
			)
			return
		}

		outPutString := ""
		db := h.db.GetDbW()

		if req.TableName == "authorized" && req.Op == "table" {
			if err := db.Exec(mysql_table.CreateAuthorizedTableSql()).Error; err != nil {
				c.AbortWithError(errno.NewError(
					http.StatusBadRequest,
					code.MySQLExecError,
					code.Text(code.MySQLExecError)+" "+err.Error()).WithErr(err),
				)
				return
			}
			outPutString = "初始化 MySQL 数据表：authorized 成功。"
		}

		if req.TableName == "authorized" && req.Op == "table_data" {
			if err := db.Exec(mysql_table.CreateAuthorizedTableDataSql()).Error; err != nil {
				c.AbortWithError(errno.NewError(
					http.StatusBadRequest,
					code.MySQLExecError,
					code.Text(code.MySQLExecError)+" "+err.Error()).WithErr(err),
				)
				return
			}
			outPutString = "初始化 MySQL 数据表：authorized 默认数据成功。"
		}

		if req.TableName == "authorized_api" && req.Op == "table" {
			if err := db.Exec(mysql_table.CreateAuthorizedAPITableSql()).Error; err != nil {
				c.AbortWithError(errno.NewError(
					http.StatusBadRequest,
					code.MySQLExecError,
					code.Text(code.MySQLExecError)+" "+err.Error()).WithErr(err),
				)
				return
			}
			outPutString = "初始化 MySQL 数据表：authorized_api 成功。"
		}

		if req.TableName == "authorized_api" && req.Op == "table_data" {
			if err := db.Exec(mysql_table.CreateAuthorizedAPITableDataSql()).Error; err != nil {
				c.AbortWithError(errno.NewError(
					http.StatusBadRequest,
					code.MySQLExecError,
					code.Text(code.MySQLExecError)+" "+err.Error()).WithErr(err),
				)
				return
			}
			outPutString = "初始化 MySQL 数据表：authorized_api 默认数据成功。"
		}

		if req.TableName == "admin" && req.Op == "table" {
			if err := db.Exec(mysql_table.CreateAdminTableSql()).Error; err != nil {
				c.AbortWithError(errno.NewError(
					http.StatusBadRequest,
					code.MySQLExecError,
					code.Text(code.MySQLExecError)+" "+err.Error()).WithErr(err),
				)
				return
			}
			outPutString = "初始化 MySQL 数据表：admin 成功。"
		}

		if req.TableName == "admin" && req.Op == "table_data" {
			if err := db.Exec(mysql_table.CreateAdminTableDataSql()).Error; err != nil {
				c.AbortWithError(errno.NewError(
					http.StatusBadRequest,
					code.MySQLExecError,
					code.Text(code.MySQLExecError)+" "+err.Error()).WithErr(err),
				)
				return
			}
			outPutString = "初始化 MySQL 数据表：admin 默认数据成功。"
		}

		if req.TableName == "menu" && req.Op == "table" {
			if err := db.Exec(mysql_table.CreateMenuTableSql()).Error; err != nil {
				c.AbortWithError(errno.NewError(
					http.StatusBadRequest,
					code.MySQLExecError,
					code.Text(code.MySQLExecError)+" "+err.Error()).WithErr(err),
				)
				return
			}
			outPutString = "初始化 MySQL 数据表：menu 成功。"
		}

		if req.TableName == "menu" && req.Op == "table_data" {
			if err := db.Exec(mysql_table.CreateMenuTableDataSql()).Error; err != nil {
				c.AbortWithError(errno.NewError(
					http.StatusBadRequest,
					code.MySQLExecError,
					code.Text(code.MySQLExecError)+" "+err.Error()).WithErr(err),
				)
				return
			}
			outPutString = "初始化 MySQL 数据表：menu 默认数据成功。"
		}

		if req.TableName == "menu_action" && req.Op == "table" {
			if err := db.Exec(mysql_table.CreateMenuActionTableSql()).Error; err != nil {
				c.AbortWithError(errno.NewError(
					http.StatusBadRequest,
					code.MySQLExecError,
					code.Text(code.MySQLExecError)+" "+err.Error()).WithErr(err),
				)
				return
			}
			outPutString = "初始化 MySQL 数据表：menu_action 成功。"
		}

		if req.TableName == "menu_action" && req.Op == "table_data" {
			if err := db.Exec(mysql_table.CreateMenuActionTableDataSql()).Error; err != nil {
				c.AbortWithError(errno.NewError(
					http.StatusBadRequest,
					code.MySQLExecError,
					code.Text(code.MySQLExecError)+" "+err.Error()).WithErr(err),
				)
				return
			}
			outPutString = "初始化 MySQL 数据表：menu_action 默认数据成功。"
		}

		if req.TableName == "admin_menu" && req.Op == "table" {
			if err := db.Exec(mysql_table.CreateAdminMenuTableSql()).Error; err != nil {
				c.AbortWithError(errno.NewError(
					http.StatusBadRequest,
					code.MySQLExecError,
					code.Text(code.MySQLExecError)+" "+err.Error()).WithErr(err),
				)
				return
			}
			outPutString = "初始化 MySQL 数据表：admin_menu 成功。"
		}

		if req.TableName == "admin_menu" && req.Op == "table_data" {
			if err := db.Exec(mysql_table.CreateAdminMenuTableDataSql()).Error; err != nil {
				c.AbortWithError(errno.NewError(
					http.StatusBadRequest,
					code.MySQLExecError,
					code.Text(code.MySQLExecError)+" "+err.Error()).WithErr(err),
				)
				return
			}
			outPutString = "初始化 MySQL 数据表：admin_menu 默认数据成功。"
		}

		c.Payload(outPutString)
	}
}
