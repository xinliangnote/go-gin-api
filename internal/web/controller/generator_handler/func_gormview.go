package generator_handler

import (
	"fmt"

	"github.com/xinliangnote/go-gin-api/configs"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"

	"go.uber.org/zap"
)

func (h *handler) GormView() core.HandlerFunc {
	return func(c core.Context) {

		type tableInfo struct {
			Name    string `db:"table_name"`    // name
			Comment string `db:"table_comment"` // comment
		}

		var tableCollect []tableInfo

		mysqlConf := configs.Get().MySQL.Read
		sqlTables := fmt.Sprintf("SELECT `table_name`,`table_comment` FROM `information_schema`.`tables` WHERE `table_schema`= '%s'", mysqlConf.Name)
		rows, err := h.db.GetDbR().Raw(sqlTables).Rows()
		if err != nil {
			h.logger.Error("rows err", zap.Error(err))

			c.HTML("generator_gorm", tableCollect)
			return
		}

		err = rows.Err()
		if err != nil {
			h.logger.Error("rows err", zap.Error(err))

			c.HTML("generator_gorm", tableCollect)
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

		c.HTML("generator_gorm", tableCollect)
	}
}
