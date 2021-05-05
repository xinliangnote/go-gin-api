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

type codes struct {
	Code    int    `json:"code"`    // 错误码
	Message string `json:"message"` // 错误码信息
}

type codeViewResponse struct {
	SystemCodes   []codes
	BusinessCodes []codes
}

func (h *handler) CodeView() core.HandlerFunc {
	return func(c core.Context) {
		fs := token.NewFileSet()
		filePath := "./internal/api/code/code.go"
		parsedFile, err := decorator.ParseFile(fs, filePath, nil, 0)
		if err != nil {
			log.Fatalf("parsing package: %s: %s\n", filePath, err)
		}

		var (
			systemCodes   []codes
			businessCodes []codes
		)

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

				if codeInt < 20000 {
					systemCodes = append(systemCodes, codes{
						Code:    codeInt,
						Message: code.Text(codeInt),
					})
				} else {
					businessCodes = append(businessCodes, codes{
						Code:    codeInt,
						Message: code.Text(codeInt),
					})
				}

			}

			return true
		})

		obj := new(codeViewResponse)
		obj.BusinessCodes = businessCodes
		obj.SystemCodes = systemCodes

		c.HTML("config_code", obj)
	}
}
