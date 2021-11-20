package upgrade

import (
	"fmt"

	"github.com/xinliangnote/go-gin-api/configs"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

type upgradeViewResponse struct {
	LockFile string            `json:"lock_file"`
	List     []upgradeViewData `json:"list"`
}

type upgradeViewData struct {
	TableName string `json:"table_name"` // 表名称
	IsHave    int32  `json:"is_have"`    // 是否已存在 1=存在 -1=不存在
}

var tableList = []string{"authorized", "authorized_api", "admin", "menu", "menu_action", "admin_menu", "cron_task"}

func (h *handler) UpgradeView() core.HandlerFunc {
	return func(c core.Context) {

		type tableInfo struct {
			Name    string // name
			Comment string // comment
		}

		var tableCollect []tableInfo

		mysqlConf := configs.Get().MySQL.Read
		sqlTables := fmt.Sprintf("SELECT `table_name`,`table_comment` FROM `information_schema`.`tables` WHERE `table_schema`= '%s'", mysqlConf.Name)
		rows, err := h.db.GetDbR().Raw(sqlTables).Rows()
		if err != nil {
			c.HTML("upgrade_view", tableCollect)
			return
		}
		defer rows.Close()

		for rows.Next() {
			var info tableInfo
			err = rows.Scan(&info.Name, &info.Comment)
			if err != nil {
				fmt.Printf("execute query tables action error,had ignored, detail is [%v]\n", err.Error())
				continue
			}

			tableCollect = append(tableCollect, info)
		}

		tableData := make([]upgradeViewData, len(tableList))
		for k, v := range tableList {
			data := upgradeViewData{
				TableName: v,
				IsHave:    -1,
			}

			tableData[k] = data
		}

		for k, v := range tableData {
			for _, haveV := range tableCollect {
				if haveV.Name == v.TableName {
					tableData[k].IsHave = 1
				}
			}
		}

		obj := new(upgradeViewResponse)
		obj.List = tableData
		obj.LockFile = configs.ProjectInstallMark
		c.HTML("upgrade_view", obj)
	}
}
