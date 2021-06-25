package install_handler

import (
	"fmt"
	"net/http"
	"os"
	"runtime"

	"github.com/xinliangnote/go-gin-api/configs"
	"github.com/xinliangnote/go-gin-api/internal/api/code"
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
	RedisAddr string `form:"redis_addr"` // 连接地址，例如：127.0.0.1:6379
	RedisPass string `form:"redis_pass"` // 连接密码
	RedisDb   string `form:"redis_db"`   // 连接 db

	MySQLAddr string `form:"mysql_addr"`
	MySQLUser string `form:"mysql_user"`
	MySQLPass string `form:"mysql_pass"`
	MySQLName string `form:"mysql_name"`
}

type sqlInfo struct {
	sqlStr   string
	tableStr string
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
				code.ConfigGoVersionError,
				code.Text(code.ConfigGoVersionError)),
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
				code.ConfigRedisConnectError,
				code.Text(code.ConfigRedisConnectError)).WithErr(err),
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
				code.ConfigMySQLConnectError,
				code.Text(code.ConfigMySQLConnectError)).WithErr(err),
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
				code.ConfigSaveError,
				code.Text(code.ConfigSaveError)).WithErr(err),
			)
			return
		}

		outPutString += "配置项 Redis、MySQL 配置成功。\n"

		createSqlInfos := []sqlInfo{
			{mysql_table.CreateAuthorizedTableSql(), "authorized"},
			{mysql_table.CreateAuthorizedTableDataSql(), "authorized"},
			{mysql_table.CreateAuthorizedAPITableSql(), "authorized_api"},
			{mysql_table.CreateAuthorizedAPITableDataSql(), "authorized_api"},
			{mysql_table.CreateAdminTableSql(), "admin"},
			{mysql_table.CreateAdminTableDataSql(), "admin"},
			{mysql_table.CreateMenuTableSql(), "menu"},
			{mysql_table.CreateMenuTableDataSql(), "menu"},
			{mysql_table.CreateMenuActionTableSql(), "menu_action"},
			{mysql_table.CreateMenuActionTableDataSql(), "menu_action"},
			{mysql_table.CreateAdminMenuTableSql(), "admin_menu"},
			{mysql_table.CreateAdminMenuTableDataSql(), "admin_menu"},
		}

		for _, info := range createSqlInfos {
			if err = db.Exec(info.sqlStr).Error; err != nil {
				c.AbortWithError(errno.NewError(
					http.StatusBadRequest,
					code.ConfigMySQLInstallError,
					"MySQL "+err.Error()).WithErr(err),
				)
				return
			}
			outPutString += "初始化 MySQL 数据表：" + info.sqlStr + " 成功。\n"
		}

		// 生成 install 完成标识
		f, err := os.Create(configs.ProjectInstallMark)
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ConfigMySQLInstallError,
				"create lock file err:  "+err.Error()).WithErr(err),
			)
			return
		}
		defer f.Close()

		c.Payload(outPutString)
	}
}
