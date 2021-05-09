package generator_handler

import (
	"bytes"
	"fmt"
	"os/exec"
	"runtime"

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

		mysqlConf := configs.Get().MySQL.Read
		shellPath := fmt.Sprintf("./scripts/gormgen.sh %s %s %s %s %s", mysqlConf.Addr, mysqlConf.User, mysqlConf.Pass, mysqlConf.Name, req.Tables)

		// runtime.GOOS = linux or darwin
		command := exec.Command("/bin/bash", "-c", shellPath)

		if runtime.GOOS == "windows" {
			command = exec.Command("cmd", "/C", shellPath)
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
