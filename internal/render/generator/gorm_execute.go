package generator_handler

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/xinliangnote/go-gin-api/configs"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

type gormExecuteRequest struct {
	Tables string `form:"tables"`
}

func (h *handler) GormExecute() core.HandlerFunc {
	return func(c core.Context) {
		req := new(gormExecuteRequest)
		if err := c.ShouldBindPostForm(req); err != nil {
			c.Payload("参数传递有误")
			return
		}

		shellPath, batPath := getCmdString(req.Tables)
		command := new(exec.Cmd)

		if runtime.GOOS == "windows" {
			command = exec.Command("cmd", "/C", batPath)
		} else {
			// runtime.GOOS = linux or darwin
			command = exec.Command("/bin/bash", "-c", shellPath)
		}

		var stderr bytes.Buffer
		command.Stderr = &stderr

		output, err := command.Output()
		if err != nil {
			c.Payload(stderr.String())
			return
		}

		c.Payload(string(output))
	}
}

func getCmdString(tables string) (string, string) {
	dir, _ := os.Getwd()
	projectPath := strings.Replace(dir, "\\", "/", -1)

	var shellPath, batPath string

	switch configs.Get().DataBaseType.Type {
	case "Mysql":
		gormgenSh := projectPath + "/scripts/gormgen.sh"
		gormgenBat := projectPath + "/scripts/gormgen.bat"
		mysqlConf := configs.Get().MySQL.Read
		shellPath = fmt.Sprintf("%s %s %s %s %s %s %s", "mysqlmd", gormgenSh, mysqlConf.Addr, mysqlConf.User, mysqlConf.Pass, mysqlConf.Name, tables)
		batPath = fmt.Sprintf("%s %s %s %s %s %s %s", "mysqlmd", gormgenBat, mysqlConf.Addr, mysqlConf.User, mysqlConf.Pass, mysqlConf.Name, tables)

	case "Postgresql":
		gormgenSh := projectPath + "/scripts/gormgen.sh"
		gormgenBat := projectPath + "/scripts/gormgen.bat"
		pgsqlConf := configs.Get().PgSQL.Read
		//shellPath = fmt.Sprintf("%s %s %s %s %s %s %s", "pgsqlcmd", gormgenSh, pgsqlConf.Addr, pgsqlConf.User, pgsqlConf.Pass, pgsqlConf.Name, tables)
		shellPath = fmt.Sprintf("%s %s %s %s %s %s %s", gormgenSh, pgsqlConf.Addr, pgsqlConf.User, pgsqlConf.Pass, pgsqlConf.Name, tables, pgsqlConf.Port)
		batPath = fmt.Sprintf("%s %s %s %s %s %s %s ", gormgenBat, pgsqlConf.Addr, pgsqlConf.User, pgsqlConf.Pass, pgsqlConf.Name, tables, pgsqlConf.Port)
	}
	return shellPath, batPath
}
