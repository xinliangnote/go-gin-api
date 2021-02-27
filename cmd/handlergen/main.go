package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"
	"unicode"
)

var handlerName string

func init() {
	handler := flag.String("handler", "", "请输入需要生成的 handler 名称\n")
	flag.Parse()

	handlerName = strings.ToLower(*handler)
}

func main() {
	fs := token.NewFileSet()
	file := fmt.Sprintf("./internal/api/controller/%s/handler.go", handlerName)
	parsedFile, err := parser.ParseFile(fs, file, nil, 0)
	if err != nil {
		log.Fatalf("parsing package: %s: %s\n", file, err)
	}

	ast.Inspect(parsedFile, func(n ast.Node) bool {
		decl, ok := n.(*ast.GenDecl)
		if !ok || decl.Tok != token.TYPE {
			return true
		}

		for _, spec := range decl.Specs {
			typeSpec, _ok := spec.(*ast.TypeSpec)
			if !_ok {
				continue
			}

			var interfaceType *ast.InterfaceType
			if interfaceType, ok = typeSpec.Type.(*ast.InterfaceType); !ok {
				continue
			}

			for _, v := range interfaceType.Methods.List {
				if len(v.Names) > 0 {
					if v.Names[0].String() == "i" {
						continue
					}

					filepath := "./internal/api/controller/" + handlerName
					filename := fmt.Sprintf("%s/func_%s.go", filepath, strings.ToLower(v.Names[0].String()))
					funcFile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0766)
					if err != nil {
						fmt.Printf("create and open func file error %v\n", err.Error())
					}

					funcContent := fmt.Sprintf("package %s\n\n", handlerName)
					funcContent += "import (\n"
					funcContent += `"github.com/xinliangnote/go-gin-api/internal/pkg/core"`
					funcContent += "\n)\n\n"
					funcContent += fmt.Sprintf("\n\ntype %sRequest struct {}\n\n", Lcfirst(v.Names[0].String()))
					funcContent += fmt.Sprintf("type %sResponse struct {}\n\n", Lcfirst(v.Names[0].String()))
					funcContent += fmt.Sprintf("func (h *handler) %s() core.HandlerFunc { \n return func(c core.Context) {\n\n}}", v.Names[0].String())

					funcFile.WriteString(funcContent)
					funcFile.Close()
				}
			}
		}
		return true
	})
}

func Lcfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}
