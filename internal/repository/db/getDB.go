package db

import (
	"github.com/xinliangnote/go-gin-api/configs"
	"github.com/xinliangnote/go-gin-api/internal/repository/iface"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql"
	"github.com/xinliangnote/go-gin-api/internal/repository/pgsql"
	"github.com/xinliangnote/go-gin-api/pkg/errors"
)

func New() (iface.Repo, error) {
	switch configs.Get().DataBaseType.Type {
	case "Mysql":
		return mysql.New()
	case "Postgresql":
		return pgsql.New()
	default:
		return nil, errors.New("不支持的数据库类型")
	}
}
