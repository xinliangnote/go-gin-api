package demo_handler

import (
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/api/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/pkg/errno"

	"github.com/pkg/errors"
)

func (h *handler) Get() core.HandlerFunc {
	type request struct {
		Name string `uri:"name"`
	}

	type response struct {
		Name string `json:"name"`
		Job  string `json:"job"`
	}

	return func(c core.Context) {
		req := new(request)
		if err := c.ShouldBindURI(req); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithErr(err),
			)
			return
		}

		if req.Name != "Tom" {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.IllegalUserName,
				code.Text(code.IllegalUserName)).WithErr(errors.New("req.Name != Tom")),
			)
			return
		}

		c.Payload(&response{
			Name: "Tom",
			Job:  "Student",
		})
	}
}
