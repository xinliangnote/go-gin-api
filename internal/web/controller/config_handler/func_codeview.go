package config_handler

import (
	"go/token"
	"log"

	"github.com/xinliangnote/go-gin-api/internal/api/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"github.com/spf13/cast"
)

func (h *handler) CodeView() core.HandlerFunc {
	return func(c core.Context) {
		fs := token.NewFileSet()
		filePath := "./internal/api/code/code.go"
		parsedFile, err := decorator.ParseFile(fs, filePath, nil, 0)
		if err != nil {
			log.Fatalf("parsing package: %s: %s\n", filePath, err)
		}

		type codes struct {
			Code    int    `json:"code"`    // 错误码
			Message string `json:"message"` // 错误码信息
		}

		var constCodes []codes

		dst.Inspect(parsedFile, func(n dst.Node) bool {
			decl, ok := n.(*dst.GenDecl)
			if !ok || decl.Tok != token.CONST {
				return true
			}

			for _, spec := range decl.Specs {
				valueSpec, _ok := spec.(*dst.ValueSpec)
				if !_ok {
					continue
				}

				codeInt := cast.ToInt(valueSpec.Values[0].(*dst.BasicLit).Value)

				constCodes = append(constCodes, codes{
					Code:    codeInt,
					Message: code.Text(codeInt),
				})
			}

			return true
		})

		c.HTML("config_code", constCodes)
	}
}
