package generator_handler

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

type handlerExecuteRequest struct {
	Name string `form:"name"`
}

func (h *handler) HandlerExecute() core.HandlerFunc {
	dir, _ := os.Getwd()
	projectPath := strings.Replace(dir, "\\", "/", -1)
	handlergenSh := projectPath + "/scripts/handlergen.sh"
	handlergenBat := projectPath + "/scripts/handlergen.bat"

	return func(c core.Context) {
		req := new(handlerExecuteRequest)
		if err := c.ShouldBindPostForm(req); err != nil {
			c.Payload("参数传递有误")
			return
		}
		shellPath := fmt.Sprintf("%s %s", handlergenSh, req.Name)
		batPath := fmt.Sprintf("%s %s", handlergenBat, req.Name)

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
