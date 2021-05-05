package install_handler

import (
	"runtime"

	"github.com/xinliangnote/go-gin-api/configs"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

type viewResponse struct {
	Config    configs.Config
	GoVersion string
}

func (h *handler) View() core.HandlerFunc {
	return func(c core.Context) {
		obj := new(viewResponse)
		obj.Config = configs.Get()
		obj.GoVersion = runtime.Version()
		c.HTML("install_view", obj)
	}
}
