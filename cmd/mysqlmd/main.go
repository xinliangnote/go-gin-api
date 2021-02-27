package main

import (
	"database/sql"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/xinliangnote/go-gin-api/configs"
	"github.com/xinliangnote/go-gin-api/internal/pkg/db"
	"github.com/xinliangnote/go-gin-api/pkg/env"
	"github.com/xinliangnote/go-gin-api/pkg/logger"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type tableInfo struct {
	Name    string         `db:"table_name"`    // name
	Comment sql.NullString `db:"table_comment"` // comment
}

type tableColumn struct {
	OrdinalPosition uint16         `db:"ORDINAL_POSITION"` // position
	ColumnName      string         `db:"COLUMN_NAME"`      // name
	ColumnType      string         `db:"COLUMN_TYPE"`      // column_type
	DataType        string         `db:"DATA_TYPE"`        // data_type
	ColumnKey       sql.NullString `db:"COLUMN_KEY"`       // key
	IsNullable      string         `db:"IS_NULLABLE"`      // nullable
	Extra           sql.NullString `db:"EXTRA"`            // extra
	ColumnComment   sql.NullString `db:"COLUMN_COMMENT"`   // comment
	ColumnDefault   sql.NullString `db:"COLUMN_DEFAULT"`   // default value
}

func main() {
	// 初始化 logger
	loggers, err := logger.NewJSONLogger(
		logger.WithField("domain", fmt.Sprintf("%s[%s]", configs.ProjectName(), env.Active().Value())),
		logger.WithTimeLayout("2006-01-02 15:04:05"),
		logger.WithFileP(configs.ProjectLogFile()),
	)
	if err != nil {
		panic(err)
	}
	defer loggers.Sync()

	// 初始化 DB
	dbRepo, err := db.New()
	if err != nil {
		loggers.Fatal("new db err", zap.Error(err))
	}

	defer func() {
		if err := dbRepo.DbWClose(); err != nil {
			loggers.Error("dbw close err", zap.Error(err))
		}

		if err := dbRepo.DbRClose(); err != nil {
			loggers.Error("dbr close err", zap.Error(err))
		}
	}()

	dbName := configs.Get().MySQL.Read.Name
	genTables := configs.Get().Cmd.GenTables
	tables, err := queryTables(dbRepo.GetDbR(), dbName, genTables)
	if err != nil {
		loggers.Error("query tables of database err", zap.Error(err))
		return
	}

	for _, table := range tables {

		filepath := "./internal/api/repository/db_repo/" + table.Name + "_repo"
		_ = os.Mkdir(filepath, 0766)

		mdName := fmt.Sprintf("%s/gen_table.md", filepath)
		mdFile, err := os.OpenFile(mdName, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0766)
		if err != nil {
			fmt.Printf("create and open markdown file error %v\n", err.Error())
			return
		}

		modelName := fmt.Sprintf("%s/gen_model.go", filepath)
		modelFile, err := os.OpenFile(modelName, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0766)
		if err != nil {
			fmt.Printf("create and open model file error %v\n", err.Error())
			return
		}

		modelContent := fmt.Sprintf("package %s%s\n", table.Name, "_repo")
		modelContent += fmt.Sprintf(`import "time"`)
		modelContent += fmt.Sprintf("\n\n// %s \n", table.Comment.String)
		modelContent += fmt.Sprintf("//go:generate gormgen -structs %s -input . \n", capitalize(table.Name))
		modelContent += fmt.Sprintf("type %s struct {\n", capitalize(table.Name))

		tableContent := fmt.Sprintf("#### %s.%s \n", dbName, table.Name)
		if table.Comment.String != "" {
			tableContent += table.Comment.String + "\n"
		}
		tableContent += "\n" +
			"| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |\n" +
			"| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |\n"

		columnInfo, columnInfoErr := queryTableColumn(dbRepo.GetDbR(), dbName, table.Name)
		if columnInfoErr != nil {
			continue
		}
		for _, info := range columnInfo {
			tableContent += fmt.Sprintf(
				"| %d | %s | %s | %s | %s | %s | %s | %s |\n",
				info.OrdinalPosition,
				info.ColumnName,
				strings.ReplaceAll(strings.ReplaceAll(info.ColumnComment.String, "|", "\\|"), "\n", ""),
				info.ColumnType,
				info.ColumnKey.String,
				info.IsNullable,
				info.Extra.String,
				info.ColumnDefault.String,
			)

			if textType(info.DataType) == "time.Time" {
				modelContent += fmt.Sprintf("%s %s `%s` // %s\n", capitalize(info.ColumnName), textType(info.DataType), "gorm:\"time\"", info.ColumnComment.String)
			} else {
				modelContent += fmt.Sprintf("%s %s // %s\n", capitalize(info.ColumnName), textType(info.DataType), info.ColumnComment.String)
			}
		}

		mdFile.WriteString(tableContent)
		mdFile.Close()

		modelContent += "}\n"
		modelFile.WriteString(modelContent)
		modelFile.Close()

	}

}

func queryTables(db *gorm.DB, dbName string, tableName string) ([]tableInfo, error) {
	var tableCollect []tableInfo
	var tableArray []string
	var commentArray []sql.NullString

	sqlTables := fmt.Sprintf("SELECT `table_name`,`table_comment` FROM `information_schema`.`tables` WHERE `table_schema`= '%s'", dbName)
	rows, err := db.Raw(sqlTables).Rows()
	if err != nil {
		return tableCollect, err
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
		tableArray = append(tableArray, info.Name)
		commentArray = append(commentArray, info.Comment)
	}

	// filter tables when specified tables params
	if tableName != "" {
		tableCollect = nil
		chooseTables := strings.Split(tableName, ",")
		indexMap := make(map[int]int)
		for _, item := range chooseTables {
			subIndexMap := getTargetIndexMap(tableArray, item)
			for k, v := range subIndexMap {
				if _, ok := indexMap[k]; ok {
					continue
				}
				indexMap[k] = v
			}
		}

		if len(indexMap) != 0 {
			for _, v := range indexMap {
				var info tableInfo
				info.Name = tableArray[v]
				info.Comment = commentArray[v]
				tableCollect = append(tableCollect, info)
			}
		}
	}

	return tableCollect, err
}

func queryTableColumn(db *gorm.DB, dbName string, tableName string) ([]tableColumn, error) {
	// 定义承载列信息的切片
	var columns []tableColumn

	sqlTableColumn := fmt.Sprintf("SELECT `ORDINAL_POSITION`,`COLUMN_NAME`,`COLUMN_TYPE`,`DATA_TYPE`,`COLUMN_KEY`,`IS_NULLABLE`,`EXTRA`,`COLUMN_COMMENT`,`COLUMN_DEFAULT` FROM `information_schema`.`columns` WHERE `table_schema`= '%s' AND `table_name`= '%s' ORDER BY `ORDINAL_POSITION` ASC",
		dbName, tableName)

	rows, err := db.Raw(sqlTableColumn).Rows()
	if err != nil {
		fmt.Printf("execute query table column action error, detail is [%v]\n", err.Error())
		return columns, err
	}
	defer rows.Close()

	for rows.Next() {
		var column tableColumn
		err = rows.Scan(
			&column.OrdinalPosition,
			&column.ColumnName,
			&column.ColumnType,
			&column.DataType,
			&column.ColumnKey,
			&column.IsNullable,
			&column.Extra,
			&column.ColumnComment,
			&column.ColumnDefault)
		if err != nil {
			fmt.Printf("query table column scan error, detail is [%v]\n", err.Error())
			return columns, err
		}
		columns = append(columns, column)
	}

	return columns, err
}

func getTargetIndexMap(tableNameArr []string, item string) map[int]int {
	indexMap := make(map[int]int)
	for i := 0; i < len(tableNameArr); i++ {
		if match, _ := regexp.MatchString(item, tableNameArr[i]); match {
			if _, ok := indexMap[i]; ok {
				continue
			}
			indexMap[i] = i
		}
	}
	return indexMap
}

func capitalize(s string) string {
	var upperStr string
	chars := strings.Split(s, "_")
	for _, val := range chars {
		vv := []rune(val)
		for i := 0; i < len(vv); i++ {
			if i == 0 {
				if vv[i] >= 97 && vv[i] <= 122 {
					vv[i] -= 32
					upperStr += string(vv[i])
				}
			} else {
				upperStr += string(vv[i])
			}
		}
	}
	return upperStr
}

func textType(s string) string {
	var mysqlTypeToGoType = map[string]string{
		"tinyint":    "int32",
		"smallint":   "int32",
		"mediumint":  "int32",
		"int":        "int32",
		"integer":    "int64",
		"bigint":     "int64",
		"float":      "float64",
		"double":     "float64",
		"decimal":    "float64",
		"date":       "string",
		"time":       "string",
		"year":       "string",
		"datetime":   "time.Time",
		"timestamp":  "time.Time",
		"char":       "string",
		"varchar":    "string",
		"tinyblob":   "string",
		"tinytext":   "string",
		"blob":       "string",
		"text":       "string",
		"mediumblob": "string",
		"mediumtext": "string",
		"longblob":   "string",
		"longtext":   "string",
	}
	return mysqlTypeToGoType[s]
}
