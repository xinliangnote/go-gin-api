package gencode_handler

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/xinliangnote/go-gin-api/configs"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

func (h *handler) InitExecute() core.HandlerFunc {
	return func(c core.Context) {
		mysqlConf := configs.Get().MySQL.Read
		shellPath := fmt.Sprintf("./scripts/init.sh %s %s %s %s", mysqlConf.Addr, mysqlConf.User, mysqlConf.Pass, mysqlConf.Name)

		command := exec.Command("/bin/bash", "-c", shellPath) //初始化 Cmd

		// windows 版本
		//command := exec.Command("cmd", "/C", shellPath)

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
