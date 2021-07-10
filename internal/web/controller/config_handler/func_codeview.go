package config_handler

import (
	"go/token"
	"log"

	"github.com/xinliangnote/go-gin-api/internal/pkg/code"
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

const (
	codeFile        = "./internal/pkg/code/code.go"
	minBusinessCode = 20000
)

func (h *handler) CodeView() core.HandlerFunc {
	fs := token.NewFileSet()
	parsedFile, err := decorator.ParseFile(fs, codeFile, nil, 0)
	if err != nil {
		log.Fatalf("parsing package: %s: %s\n", codeFile, err)
	}

	var (
		systemCodes   []codes
		businessCodes []codes
	)

	dst.Inspect(parsedFile, func(n dst.Node) bool {
		// GenDecl 代表除函数以外的所有声明，即 import、const、var 和 type
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

			if codeInt < minBusinessCode {
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

	return func(c core.Context) {
		obj := new(codeViewResponse)
		obj.BusinessCodes = businessCodes
		obj.SystemCodes = systemCodes

		c.HTML("config_code", obj)
	}
}
