package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/xinliangnote/go-gin-api/cmd/pgsqlcmd/pgsql"

	"gorm.io/gorm"
)

type tableInfo struct {
	Name    string         `db:"table_name"`    // name
	Comment sql.NullString `db:"table_comment"` // comment
}

type tableColumn struct {
	OrdinalPosition uint16 `db:"COLUMN_NUMBER"` // position
	ColumnName      string `db:"COLUMN_NAME"`   // name
	ColumnType      string `db:"COLUMN_TYPE"`   // column_type
	//DataType        string         `db:"DATA_TYPE"`     // data_type
	ColumnKey  sql.NullString `db:"COLUMN_KEY"`  // key
	IsNullable string         `db:"IS_NULLABLE"` // nullable
	//Extra           sql.NullString `db:"EXTRA"`            // extra
	ColumnComment sql.NullString `db:"COLUMN_COMMENT"` // comment
	ColumnDefault sql.NullString `db:"COLUMN_DEFAULT"` // default value
}

var (
	dbAddr    string
	dbUser    string
	dbPass    string
	dbPort    string
	dbName    string
	genTables string
)

func init() {
	addr := flag.String("addr", "127.0.0.1", "请输入 db 地址，例如：127.0.0.1:3306\n")
	user := flag.String("user", "postgres", "请输入 db 用户名\n")
	pass := flag.String("pass", "123456", "请输入 db 密码\n")
	name := flag.String("name", "mydb", "请输入 db 名称\n")
	port := flag.String("port", "5432", "请输入 db 端口\n")
	table := flag.String("tables", "cc", "请输入 table 名称，默认为“*”，多个可用“,”分割\n")

	flag.Parse()

	dbAddr = *addr
	dbUser = *user
	dbPass = *pass
	dbPort = *port
	dbName = strings.ToLower(*name)
	genTables = strings.ToLower(*table)
}

func main() {
	// 初始化 DB
	db, err := pgsql.New(dbAddr, dbUser, dbPass, dbName, dbPort)
	if err != nil {
		log.Fatal("new db err", err)
	}

	defer func() {
		if err := db.DbClose(); err != nil {
			log.Println("db close err", err)
		}
	}()

	tables, err := queryTables(db.GetDb(), genTables)
	if err != nil {
		log.Println("query tables of database err", err)
		return
	}

	for _, table := range tables {

		filepath := "./internal/repository/pgsql/" + table.Name
		_ = os.Mkdir(filepath, 0766)
		fmt.Println("create dir : ", filepath)

		mdName := fmt.Sprintf("%s/gen_table.md", filepath)
		mdFile, err := os.OpenFile(mdName, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0766)
		if err != nil {
			fmt.Printf("markdown file error %v\n", err.Error())
			return
		}
		fmt.Println("  └── file : ", table.Name+"/gen_table.md")

		modelName := fmt.Sprintf("%s/gen_model.go", filepath)
		modelFile, err := os.OpenFile(modelName, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0766)
		if err != nil {
			fmt.Printf("create and open model file error %v\n", err.Error())
			return
		}
		fmt.Println("  └── file : ", table.Name+"/gen_model.go")

		modelContent := fmt.Sprintf("package %s\n", table.Name)
		modelContent += fmt.Sprintf(`import "time"`)
		modelContent += fmt.Sprintf("\n\n// %s %s \n", capitalize(table.Name), table.Comment.String)
		modelContent += fmt.Sprintf("//go:generate gormgen -structs %s -input . \n", capitalize(table.Name))
		modelContent += fmt.Sprintf("type %s struct {\n", capitalize(table.Name))

		tableContent := fmt.Sprintf("#### %s.%s \n", dbName, table.Name)
		if table.Comment.String != "" {
			tableContent += table.Comment.String + "\n"
		}
		tableContent += "\n" +
			"| 序号 | 名称 | 描述 | 类型 | 是否主键 | 是否为空  | 默认值 |\n" +
			"| :--: | :--: | :--: | :--: | :--:  | :--: | :--: |\n"

		columnInfo, columnInfoErr := queryTableColumn(db.GetDb(), table.Name)
		if columnInfoErr != nil {
			continue
		}
		for _, info := range columnInfo {
			tableContent += fmt.Sprintf(
				"| %d | %s | %s | %s | %s | %s | %s  |\n",
				info.OrdinalPosition,
				info.ColumnName,
				strings.ReplaceAll(strings.ReplaceAll(info.ColumnComment.String, "|", "\\|"), "\n", ""),
				textType(info.ColumnType),
				info.ColumnKey.String,
				info.IsNullable,
				//info.Extra.String,
				info.ColumnDefault.String,
			)

			if textType(info.ColumnType) == "time.Time" {
				modelContent += fmt.Sprintf("%s %s `%s` // %s\n", capitalize(info.ColumnName), textType(info.ColumnType), "gorm:\"time\"", info.ColumnComment.String)
			} else {
				modelContent += fmt.Sprintf("%s %s // %s\n", capitalize(info.ColumnName), textType(info.ColumnType), info.ColumnComment.String)
			}
		}

		_, _ = mdFile.WriteString(tableContent)
		_ = mdFile.Close()

		modelContent += "}\n"
		_, _ = modelFile.WriteString(modelContent)
		_ = modelFile.Close()

	}

}

func queryTables(db *gorm.DB, tableName string) ([]tableInfo, error) {
	var tableCollect []tableInfo
	var tableArray []string
	var commentArray []sql.NullString

	//sqlTables := fmt.Sprintf("SELECT `table_name`,`table_comment` FROM `information_schema`.`tables` WHERE `table_schema`= '%s'", dbName)
	sqlTables := `
			select
				relname as tab_name,
				obj_description(c.oid) as table_comment
			from
				pg_class c
			where
				obj_description(c.oid) is not null 

		`

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
	if tableName != "*" {
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

func queryTableColumn(db *gorm.DB, tableName string) ([]tableColumn, error) {
	// 定义承载列信息的切片
	var columns []tableColumn

	//sqlTableColumn := fmt.Sprintf("SELECT `ORDINAL_POSITION`,`COLUMN_NAME`,`COLUMN_TYPE`,`DATA_TYPE`,`COLUMN_KEY`,`IS_NULLABLE`,`EXTRA`,`COLUMN_COMMENT`,`COLUMN_DEFAULT` FROM `information_schema`.`columns` WHERE `table_schema`= '%s' AND `table_name`= '%s' ORDER BY `ORDINAL_POSITION` ASC",
	//	dbName, tableName)
	sqlTableColumn := fmt.Sprintf(`
			SELECT
	           a.attnum AS COLUMN_NUMBER,
	           a.attname AS COLUMN_NAME,
	           (case  when  d.description is NULL then ' '  when d.description is Not NULl then  d.description end ) as COLUMN_COMMENT,
	           --format_type(a.atttypid, a.atttypmod) AS COLUMN_TYPE,
	           a.attnotnull AS IS_NULLABLE,
				COALESCE(pg_get_expr(ad.adbin, ad.adrelid), '') AS COLUMN_DEFAULT,
	   		COALESCE(ct.contype = 'p', false) AS  COLUMN_KEY,
	   		CASE
	       	WHEN a.atttypid = ANY ('{int,int8,int2}'::regtype[])
	         		AND EXISTS (
					SELECT 1 FROM pg_attrdef ad
	            	WHERE  ad.adrelid = a.attrelid
	            	AND    ad.adnum   = a.attnum
	            	-- AND    ad.adsrc = 'nextval('''
	               --	|| (pg_get_serial_sequence (a.attrelid::regclass::text
	               --	                          , a.attname))::regclass
	               --	|| '''::regclass)'
	            	)
	           THEN CASE a.atttypid
	                   WHEN 'int'::regtype  THEN 'serial'
	                   WHEN 'int8'::regtype THEN 'bigserial'
	                   WHEN 'int2'::regtype THEN 'smallserial'
	                END
				WHEN a.atttypid = ANY ('{uuid}'::regtype[]) AND COALESCE(pg_get_expr(ad.adbin, ad.adrelid), '') != ''
	           THEN 'autogenuuid'
	       	ELSE format_type(a.atttypid, a.atttypmod)
	   		END AS COLUMN_TYPE
			FROM pg_attribute a
			JOIN ONLY pg_class c ON c.oid = a.attrelid
			JOIN ONLY pg_namespace n ON n.oid = c.relnamespace
			LEFT JOIN pg_constraint ct ON ct.conrelid = c.oid
			AND a.attnum = ANY(ct.conkey) AND ct.contype = 'p'
			left join pg_description d on d.objoid = a.attrelid
			and d.objsubid = a.attnum
			LEFT JOIN pg_attrdef ad ON ad.adrelid = c.oid AND ad.adnum = a.attnum
			WHERE a.attisdropped = false
			AND n.nspname = 'public'
			AND c.relname = '%v'
			AND a.attnum > 0
			ORDER BY a.attnum
		`, tableName)

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
			&column.ColumnComment,
			&column.IsNullable,
			&column.ColumnDefault,
			//&column.DataType,
			&column.ColumnKey,
			&column.ColumnType,

			//&column.Extra,

		)
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
		if tableNameArr[i] == item {
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
				}
				upperStr += string(vv[i])
			} else {
				upperStr += string(vv[i])
			}
		}
	}
	return upperStr
}

func textType(s string) string {
	if strings.Contains(s, "char") || in(s, []string{
		"text",
	}) {
		return "string"
	}
	// postgres
	{
		if in(s, []string{"double precision", "double"}) {
			return "float64"
		}
		if in(s, []string{"bigint", "bigserial", "big serial"}) {
			return "int64"
		}
		if in(s, []string{"integer", "smallint", "serial", "smallserial"}) {
			return "int32"
		}
		if in(s, []string{"numeric", "decimal", "real"}) {
			return "decimal.Decimal"
		}
		if in(s, []string{"bytea"}) {
			return "[]byte"
		}
		if strings.Contains(s, "time") || in(s, []string{"date", "datetime", "timestamp"}) {
			return "time.Time"
		}
		if in(s, []string{"jsonb"}) {
			return "json.RawMessage"
		}
		if in(s, []string{"bool", "boolean"}) {
			return "bool"
		}

		if in(s, []string{"bigint[]"}) {
			return "[]int64"
		}
		if in(s, []string{"integer[]"}) {
			return "[]int64"
		}
		if in(s, []string{"text[]"}) {
			return "pq.StringArray"

		}
	}
	// mysql
	{
		if strings.HasPrefix(s, "int") {
			return "int32"
		}
		if strings.HasPrefix(s, "varchar") {
			return "string"
		}
		if s == "json" {
			return "json.RawMessage"
		}
		if in(s, []string{"bool", "boolean"}) {
			return "bool"
		}
	}

	return s
}

// s 是否in arr
func in(s string, arr []string) bool {
	for _, v := range arr {
		if v == s {
			return true
		}
	}
	return false
}

//func textType(s string) string {
//	var pgsqlTypeToGoType = map[string]string{
//		"tinyint":    "int32",
//		"smallint":   "int32",
//		"mediumint":  "int32",
//		"int":        "int32",
//		"integer":    "int64",
//		"bigint":     "int64",
//		"float":      "float64",
//		"double":     "float64",
//		"decimal":    "float64",
//		"date":       "string",
//		"time":       "string",
//		"year":       "string",
//		"datetime":   "time.Time",
//		"timestamp":  "time.Time",
//		"char":       "string",
//		"varchar":    "string",
//		"tinyblob":   "string",
//		"tinytext":   "string",
//		"blob":       "string",
//		"text":       "string",
//		"mediumblob": "string",
//		"mediumtext": "string",
//		"longblob":   "string",
//		"longtext":   "string",
//	}
//	return pgsqlTypeToGoType[s]
//}
