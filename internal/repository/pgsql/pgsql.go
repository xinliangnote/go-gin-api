package pgsql

import (
	"fmt"
	"time"

	"github.com/xinliangnote/go-gin-api/configs"
	"github.com/xinliangnote/go-gin-api/internal/repository/iface"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql"
	"github.com/xinliangnote/go-gin-api/pkg/errors"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	//"gorm.io/gorm/schema"
)

// Predicate is a string that acts as a condition in the where clause
type Predicate string

var (
	EqualPredicate              = Predicate("=")
	NotEqualPredicate           = Predicate("<>")
	GreaterThanPredicate        = Predicate(">")
	GreaterThanOrEqualPredicate = Predicate(">=")
	SmallerThanPredicate        = Predicate("<")
	SmallerThanOrEqualPredicate = Predicate("<=")
	LikePredicate               = Predicate("LIKE")
)

type dbRepo struct {
	DbR *gorm.DB
	DbW *gorm.DB
}

func New() (iface.Repo, error) {
	cfg := configs.Get().PgSQL
	dbr, err := dbConnect(cfg.Read.User, cfg.Read.Pass, cfg.Read.Addr, cfg.Read.Name, cfg.Read.Port)
	if err != nil {
		return nil, err
	}

	dbw, err := dbConnect(cfg.Write.User, cfg.Write.Pass, cfg.Write.Addr, cfg.Write.Name, cfg.Read.Port)
	if err != nil {
		return nil, err
	}

	return &dbRepo{
		DbR: dbr,
		DbW: dbw,
	}, nil
}

func (d *dbRepo) i() {}

func (d *dbRepo) GetDbR() *gorm.DB {
	return d.DbR
}

func (d *dbRepo) GetDbW() *gorm.DB {
	return d.DbW
}

func (d *dbRepo) DbRClose() error {
	sqlDB, err := d.DbR.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (d *dbRepo) DbWClose() error {
	sqlDB, err := d.DbW.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func dbConnect(user, pass, addr, dbName, port string) (*gorm.DB, error) {
	//	dataSouce := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s", "localhost", "5432", "postgres", "game", "disable", "123")
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
		addr,
		port,
		user,
		dbName,
		"disable",
		pass,
	)
	//gorm.Open("postgres", dataSource)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		//Logger: logger.Default.LogMode(logger.Info), // 日志配置
	})

	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("[db connection failed] Database name: %s", dbName))
	}

	db.Set("gorm:table_options", "CHARSET=utf8mb4")

	cfg := configs.Get().PgSQL.Base

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// 设置连接池 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConn)

	// 设置最大连接数 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConn)

	// 设置最大连接超时
	sqlDB.SetConnMaxLifetime(time.Minute * cfg.ConnMaxLifeTime)

	// 使用插件
	db.Use(&mysql.TracePlugin{})

	return db, nil
}
