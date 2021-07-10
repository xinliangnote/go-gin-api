package install_handler

import (
	"fmt"
	"net/http"
	"os"
	"runtime"

	"github.com/xinliangnote/go-gin-api/configs"
	"github.com/xinliangnote/go-gin-api/internal/pkg/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/web/controller/install_handler/mysql_table"
	"github.com/xinliangnote/go-gin-api/pkg/env"
	"github.com/xinliangnote/go-gin-api/pkg/errno"

	"github.com/go-redis/redis/v7"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type initExecuteRequest struct {
	Language  string `form:"language" `  // 语言包
	RedisAddr string `form:"redis_addr"` // 连接地址，例如：127.0.0.1:6379
	RedisPass string `form:"redis_pass"` // 连接密码
	RedisDb   string `form:"redis_db"`   // 连接 db

	MySQLAddr string `form:"mysql_addr"`
	MySQLUser string `form:"mysql_user"`
	MySQLPass string `form:"mysql_pass"`
	MySQLName string `form:"mysql_name"`
}

func (h *handler) Execute() core.HandlerFunc {
	return func(c core.Context) {
		req := new(initExecuteRequest)
		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithErr(err),
			)
			return
		}

		versionStr := runtime.Version()
		version := cast.ToFloat32(versionStr[2:6])
		if version < 1.15 {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.GoVersionError,
				code.Text(code.GoVersionError)),
			)
			return
		}

		outPutString := ""

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
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.RedisConnectError,
				code.Text(code.RedisConnectError)).WithErr(err),
			)
			return
		}

		defer redisClient.Close()

		outPutString += "已检测 Redis 配置可用。\n"

		dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
			req.MySQLUser,
			req.MySQLPass,
			req.MySQLAddr,
			req.MySQLName,
			true,
			"Local")

		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
			//Logger: logger.Default.LogMode(logger.Info), // 日志配置
		})

		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.MySQLConnectError,
				code.Text(code.MySQLConnectError)).WithErr(err),
			)
			return
		}

		db.Set("gorm:table_options", "CHARSET=utf8mb4")

		dbClient, _ := db.DB()
		defer dbClient.Close()

		outPutString += "已检测 MySQL 配置可用。\n"

		viper.SetConfigName(env.Active().Value() + "_configs")
		viper.SetConfigType("toml")
		viper.AddConfigPath("./configs")

		viper.Set("language.local", req.Language)

		viper.Set("redis.addr", req.RedisAddr)
		viper.Set("redis.pass", req.RedisPass)
		viper.Set("redis.db", req.RedisDb)

		viper.Set("mysql.read.addr", req.MySQLAddr)
		viper.Set("mysql.read.user", req.MySQLUser)
		viper.Set("mysql.read.pass", req.MySQLPass)
		viper.Set("mysql.read.name", req.MySQLName)

		viper.Set("mysql.write.addr", req.MySQLAddr)
		viper.Set("mysql.write.user", req.MySQLUser)
		viper.Set("mysql.write.pass", req.MySQLPass)
		viper.Set("mysql.write.name", req.MySQLName)

		if viper.WriteConfig() != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.WriteConfigError,
				code.Text(code.WriteConfigError)).WithErr(err),
			)
			return
		}

		outPutString += "语言包 " + req.Language + " 配置成功。\n"
		outPutString += "配置项 Redis、MySQL 配置成功。\n"

		if err = db.Exec(mysql_table.CreateAuthorizedTableSql()).Error; err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.MySQLExecError,
				code.Text(code.MySQLExecError)+" "+err.Error()).WithErr(err),
			)
			return
		}
		outPutString += "初始化 MySQL 数据表：authorized 成功。\n"

		if err = db.Exec(mysql_table.CreateAuthorizedTableDataSql()).Error; err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.MySQLExecError,
				code.Text(code.MySQLExecError)+" "+err.Error()).WithErr(err),
			)
			return
		}
		outPutString += "初始化 MySQL 数据表：authorized 默认数据成功。\n"

		if err = db.Exec(mysql_table.CreateAuthorizedAPITableSql()).Error; err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.MySQLExecError,
				code.Text(code.MySQLExecError)+" "+err.Error()).WithErr(err),
			)
			return
		}
		outPutString += "初始化 MySQL 数据表：authorized_api 成功。\n"

		if err = db.Exec(mysql_table.CreateAuthorizedAPITableDataSql()).Error; err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.MySQLExecError,
				code.Text(code.MySQLExecError)+" "+err.Error()).WithErr(err),
			)
			return
		}
		outPutString += "初始化 MySQL 数据表：authorized_api 默认数据成功。\n"

		if err = db.Exec(mysql_table.CreateAdminTableSql()).Error; err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.MySQLExecError,
				code.Text(code.MySQLExecError)+" "+err.Error()).WithErr(err),
			)
			return
		}
		outPutString += "初始化 MySQL 数据表：admin 成功。\n"

		if err = db.Exec(mysql_table.CreateAdminTableDataSql()).Error; err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.MySQLExecError,
				code.Text(code.MySQLExecError)+" "+err.Error()).WithErr(err),
			)
			return
		}
		outPutString += "初始化 MySQL 数据表：admin 默认数据成功。\n"

		if err = db.Exec(mysql_table.CreateMenuTableSql()).Error; err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.MySQLExecError,
				code.Text(code.MySQLExecError)+" "+err.Error()).WithErr(err),
			)
			return
		}
		outPutString += "初始化 MySQL 数据表：menu 成功。\n"

		if err = db.Exec(mysql_table.CreateMenuTableDataSql()).Error; err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.MySQLExecError,
				code.Text(code.MySQLExecError)+" "+err.Error()).WithErr(err),
			)
			return
		}
		outPutString += "初始化 MySQL 数据表：menu 默认数据成功。\n"

		if err = db.Exec(mysql_table.CreateMenuActionTableSql()).Error; err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.MySQLExecError,
				code.Text(code.MySQLExecError)+" "+err.Error()).WithErr(err),
			)
			return
		}
		outPutString += "初始化 MySQL 数据表：menu_action 成功。\n"

		if err = db.Exec(mysql_table.CreateMenuActionTableDataSql()).Error; err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.MySQLExecError,
				code.Text(code.MySQLExecError)+" "+err.Error()).WithErr(err),
			)
			return
		}
		outPutString += "初始化 MySQL 数据表：menu_action 默认数据成功。\n"

		if err = db.Exec(mysql_table.CreateAdminMenuTableSql()).Error; err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.MySQLExecError,
				code.Text(code.MySQLExecError)+" "+err.Error()).WithErr(err),
			)
			return
		}
		outPutString += "初始化 MySQL 数据表：admin_menu 成功。\n"

		if err = db.Exec(mysql_table.CreateAdminMenuTableDataSql()).Error; err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.MySQLExecError,
				code.Text(code.MySQLExecError)+" "+err.Error()).WithErr(err),
			)
			return
		}
		outPutString += "初始化 MySQL 数据表：admin_menu 默认数据成功。\n"

		// 生成 install 完成标识
		f, err := os.Create(configs.ProjectInstallMark)
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.MySQLExecError,
				code.Text(code.MySQLExecError)+" "+err.Error()).WithErr(err),
			)
			return
		}
		defer f.Close()

		c.Payload(outPutString)
	}
}
