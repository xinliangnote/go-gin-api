package gencode_handler

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

type handlerExecuteRequest struct {
	Name string `form:"name"`
}

func (h *handler) HandlerExecute() core.HandlerFunc {
	return func(c core.Context) {
		req := new(handlerExecuteRequest)
		if err := c.ShouldBindPostForm(req); err != nil {
			c.Payload("参数传递有误")
			return
		}

		shellPath := fmt.Sprintf("./scripts/handlergen.sh %s", req.Name)
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
