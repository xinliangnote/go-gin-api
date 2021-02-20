package demo_handler

import (
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/api/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/pkg/errno"

	"github.com/pkg/errors"
)

func (h *handler) Post() core.HandlerFunc {
	type request struct {
		Name string `form:"name"`
	}

	type response struct {
		Name string `json:"name"`
		Job  string `json:"job"`
	}

	return func(c core.Context) {
		req := new(request)
		if err := c.ShouldBindPostForm(req); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithErr(err),
			)
			return
		}

		if req.Name != "Jack" {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.IllegalUserName,
				code.Text(code.IllegalUserName)).WithErr(errors.New("req.Name != Jack")),
			)
			return
		}

		c.Payload(&response{
			Name: "Jack",
			Job:  "Teacher",
		})
	}
}
