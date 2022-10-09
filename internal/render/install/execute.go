package install

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
	"strings"

	"github.com/xinliangnote/go-gin-api/configs"
	"github.com/xinliangnote/go-gin-api/internal/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/proposal/tablesqls"
	"github.com/xinliangnote/go-gin-api/pkg/errors"

	"github.com/go-redis/redis/v7"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type initExecuteRequest struct {
	Language  string `form:"language" `  // 语言包
	RedisAddr string `form:"redis_addr"` // 连接地址，例如：127.0.0.1:6379
	RedisPass string `form:"redis_pass"` // 连接密码
	RedisDb   string `form:"redis_db"`   // 连接 db

	DataBaseType string `form:"database_type"` //连接数据库类型

	DataBaseAddr string `form:"database_addr"`
	DataBaseUser string `form:"database_user"`
	DataBasePass string `form:"database_pass"`
	DataBaseName string `form:"database_name"`
}

func (h *handler) Execute() core.HandlerFunc {

	dateBaseTypeErrCode := map[string]int{
		"Mysql":      code.MySQLConnectError,
		"Postgresql": code.PgSQlConnectError,
	}
	dataBaseKey := map[string]string{
		"Mysql":      "table_sql",
		"Postgresql": "table_pgsql",
	}
	dataKey := map[string]string{
		"Mysql":      "table_data_sql",
		"Postgresql": "table_data_pgsql",
	}

	installTableList := map[string]map[string]string{
		"authorized": {
			"table_sql":        tablesqls.CreateAuthorizedTableSql(),
			"table_pgsql":      tablesqls.CreateAuthorizedTablePGSql(),
			"table_data_sql":   tablesqls.CreateAuthorizedTableDataSql(),
			"table_data_pgsql": tablesqls.CreateAuthorizedTableDataPGSql(),
		},
		"authorized_api": {
			"table_sql":        tablesqls.CreateAuthorizedAPITableSql(),
			"table_pgsql":      tablesqls.CreateAuthorizedAPITablePGSql(),
			"table_data_sql":   tablesqls.CreateAuthorizedAPITableDataSql(),
			"table_data_pgsql": tablesqls.CreateAuthorizedAPITableDataPGSql(),
		},
		"admin": {
			"table_sql":        tablesqls.CreateAdminTableSql(),
			"table_pgsql":      tablesqls.CreateAdminTablePGSql(),
			"table_data_sql":   tablesqls.CreateAdminTableDataSql(),
			"table_data_pgsql": tablesqls.CreateAdminTableDataPGSql(),
		},
		"admin_menu": {
			"table_sql":        tablesqls.CreateAdminMenuTableSql(),
			"table_pgsql":      tablesqls.CreateAdminMenuTablePGSql(),
			"table_data_sql":   tablesqls.CreateAdminMenuTableDataSql(),
			"table_data_pgsql": tablesqls.CreateAdminMenuTableDataPGSql(),
		},
		"menu": {
			"table_sql":        tablesqls.CreateMenuTableSql(),
			"table_pgsql":      tablesqls.CreateMenuTablePGSql(),
			"table_data_sql":   tablesqls.CreateMenuTableDataSql(),
			"table_data_pgsql": tablesqls.CreateMenuTableDataPGSql(),
		},
		"menu_action": {
			"table_sql":        tablesqls.CreateMenuActionTableSql(),
			"table_pgsql":      tablesqls.CreateMenuActionTablePGSql(),
			"table_data_sql":   tablesqls.CreateMenuActionTableDataSql(),
			"table_data_pgsql": tablesqls.CreateMenuActionTableDataPGSql(),
		},
		"cron_task": {
			"table_sql":        tablesqls.CreateCronTaskTableSql(),
			"table_pgsql":      tablesqls.CreateCronTaskTablePGSql(),
			"table_data_sql":   "",
			"table_data_pgsql": "",
		},
	}

	return func(ctx core.Context) {
		req := new(initExecuteRequest)
		if err := ctx.ShouldBindForm(req); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}

		// region 验证 version
		versionStr := runtime.Version()
		version := cast.ToFloat32(versionStr[2:6])
		if version < configs.MinGoVersion {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.GoVersionError,
				code.Text(code.GoVersionError)),
			)
			return
		}
		// endregion

		// region 验证 Redis 配置
		cfg := configs.Get()
		redisClient := redis.NewClient(&redis.Options{
			Addr:         req.RedisAddr,
			Password:     req.RedisPass,
			DB:           cast.ToInt(req.RedisDb),
			MaxRetries:   cfg.Redis.MaxRetries,
			PoolSize:     cfg.Redis.PoolSize,
			MinIdleConns: cfg.Redis.MinIdleConns,
		})

		if err := redisClient.Ping().Err(); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.RedisConnectError,
				code.Text(code.RedisConnectError)).WithError(err),
			)
			return
		}

		defer redisClient.Close()

		outPutString := "已检测 Redis 配置可用。\n"
		// endregion

		// region 验证数据库配置
		db, err := getDB(*req)
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				dateBaseTypeErrCode[req.DataBaseType],
				code.Text(dateBaseTypeErrCode[req.DataBaseType])).WithError(err),
			)
			return
		}
		db.Set("gorm:table_options", "CHARSET=utf8mb4")

		dbClient, _ := db.DB()
		defer dbClient.Close()

		outPutString += "已检测 " + req.DataBaseType + " 配置可用。\n"
		// endregion

		// region 写入配置文件
		viper.Set("language.local", req.Language)

		viper.Set("redis.addr", req.RedisAddr)
		viper.Set("redis.pass", req.RedisPass)
		viper.Set("redis.db", req.RedisDb)

		viper.Set("mysql.read.addr", req.DataBaseAddr)
		viper.Set("mysql.read.user", req.DataBaseUser)
		viper.Set("mysql.read.pass", req.DataBasePass)
		viper.Set("mysql.read.name", req.DataBaseName)

		viper.Set("mysql.write.addr", req.DataBaseAddr)
		viper.Set("mysql.write.user", req.DataBaseUser)
		viper.Set("mysql.write.pass", req.DataBasePass)
		viper.Set("mysql.write.name", req.DataBaseName)

		viper.Set("pgsql.read.addr", strings.Split(req.DataBaseAddr, ":")[0])
		viper.Set("pgsql.read.user", req.DataBaseUser)
		viper.Set("pgsql.read.pass", req.DataBasePass)
		viper.Set("pgsql.read.name", req.DataBaseName)
		viper.Set("pgsql.read.port", strings.Split(req.DataBaseAddr, ":")[1])

		viper.Set("pgsql.write.addr", strings.Split(req.DataBaseAddr, ":")[0])
		viper.Set("pgsql.write.user", req.DataBaseUser)
		viper.Set("pgsql.write.pass", req.DataBasePass)
		viper.Set("pgsql.write.name", req.DataBaseName)
		viper.Set("pgsql.write.port", strings.Split(req.DataBaseAddr, ":")[1])

		viper.Set("databasetype.type", req.DataBaseType)

		if viper.WriteConfig() != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.WriteConfigError,
				code.Text(code.WriteConfigError)).WithError(err),
			)
			return
		}

		outPutString += "语言包 " + req.Language + " 配置成功。\n"
		outPutString += "配置项 Redis、" + req.DataBaseType + " 配置成功。\n"
		// endregion

		// region 初始化表结构 + 默认数据
		for k, v := range installTableList {
			if v[dataBaseKey[req.DataBaseType]] != "" {
				// region 初始化表结构
				if err = db.Exec(v[dataBaseKey[req.DataBaseType]]).Error; err != nil {
					ctx.AbortWithError(core.Error(
						http.StatusBadRequest,
						code.MySQLExecError,
						code.Text(code.MySQLExecError)+" "+err.Error()).WithError(err),
					)
					return
				}

				outPutString += "初始化 " + req.DataBaseType + " 数据表：" + k + " 成功。\n"
				// endregion

				// region 初始化默认数据
				if v[dataKey[req.DataBaseType]] != "" {
					if err = db.Exec(v[dataKey[req.DataBaseType]]).Error; err != nil {
						ctx.AbortWithError(core.Error(
							http.StatusBadRequest,
							code.MySQLExecError,
							code.Text(code.MySQLExecError)+" "+err.Error()).WithError(err),
						)
						return
					}

					outPutString += "初始化 " + req.DataBaseType + " 数据表：" + k + " 默认数据成功。\n"
				}
				// endregion
			}

		}
		// endregion

		// region 生成 install 完成标识
		f, err := os.Create(configs.ProjectInstallMark)
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.MySQLExecError,
				code.Text(code.MySQLExecError)+" "+err.Error()).WithError(err),
			)
			return
		}
		defer f.Close()
		// endregion

		ctx.Payload(outPutString)
	}
}

func getDB(req initExecuteRequest) (*gorm.DB, error) {
	switch req.DataBaseType {
	case "Mysql":
		// region 验证 MySQL 配置
		dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
			req.DataBaseUser,
			req.DataBasePass,
			req.DataBaseAddr,
			req.DataBaseName,
			true,
			"Local")
		return gorm.Open(mysql.Open(dsn), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
			//Logger: logger.Default.LogMode(logger.Info), // 日志配置
		})
	case "Postgresql":
		// region 验证 PgsqlSQL 配置
		pgDsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
			strings.Split(req.DataBaseAddr, ":")[0],
			strings.Split(req.DataBaseAddr, ":")[1],
			//req.PgSQLPort,
			req.DataBaseUser,
			req.DataBaseName,
			"disable",
			req.DataBasePass,
		)
		//gorm.Open("postgres", dataSource)
		return gorm.Open(postgres.Open(pgDsn), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
			//Logger: logger.Default.LogMode(logger.Info), // 日志配置
		})
	default:
		return nil, errors.New("暂不支持数据库类型")
	}
}
