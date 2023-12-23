package pgsql

import (
	"fmt"

	"github.com/xinliangnote/go-gin-api/pkg/errors"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var _ Repo = (*dbRepo)(nil)

type Repo interface {
	i()
	GetDb() *gorm.DB
	DbClose() error
}

type dbRepo struct {
	DbConn *gorm.DB
}

func New(dbAddr, dbUser, dbPass, dbName, port string) (Repo, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
		dbAddr,
		port,
		dbUser,
		dbName,
		"disable",
		dbPass,
	)

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

	return &dbRepo{
		DbConn: db,
	}, nil
}

func (d *dbRepo) i() {}

func (d *dbRepo) GetDb() *gorm.DB {
	return d.DbConn
}

func (d *dbRepo) DbClose() error {
	sqlDB, err := d.DbConn.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
