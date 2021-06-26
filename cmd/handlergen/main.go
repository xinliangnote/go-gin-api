package main

import (
	"flag"
	"fmt"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"unicode"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
)

var handlerName string

func init() {
	handler := flag.String("handler", "", "请输入需要生成的 handler 名称\n")
	flag.Parse()

	handlerName = strings.ToLower(*handler)
}

func main() {
	fs := token.NewFileSet()
	filePath := fmt.Sprintf("./internal/api/controller/%s", handlerName)
	parsedFile, err := decorator.ParseFile(fs, filePath+"/handler.go", nil, 0)
	if err != nil {
		log.Fatalf("parsing package: %s: %s\n", filePath, err)
	}

	files, _ := ioutil.ReadDir(filePath)
	if len(files) > 1 {
		log.Fatalf("请先确保 %s 目录中，有且仅有 handler.go 一个文件。", filePath)
	}

	dst.Inspect(parsedFile, func(n dst.Node) bool {
		decl, ok := n.(*dst.GenDecl)
		if !ok || decl.Tok != token.TYPE {
			return true
		}

		for _, spec := range decl.Specs {
			typeSpec, _ok := spec.(*dst.TypeSpec)
			if !_ok {
				continue
			}

			var interfaceType *dst.InterfaceType
			if interfaceType, ok = typeSpec.Type.(*dst.InterfaceType); !ok {
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
						continue
					}

					if funcFile == nil {
						fmt.Printf("func file is nil \n")
						continue
					}

					fmt.Println("  └── file : ", filename)

					funcContent := fmt.Sprintf("package %s\n\n", handlerName)
					funcContent += "import (\n"
					funcContent += `"github.com/xinliangnote/go-gin-api/internal/pkg/core"`
					funcContent += "\n)\n\n"
					funcContent += fmt.Sprintf("\n\ntype %sRequest struct {}\n\n", Lcfirst(v.Names[0].String()))
					funcContent += fmt.Sprintf("type %sResponse struct {}\n\n", Lcfirst(v.Names[0].String()))

					// 首行注释
					funcContent += fmt.Sprintf("%s\n", v.Decorations().Start.All()[0])

					nameArr := strings.Split(v.Decorations().Start.All()[0], v.Names[0].String())
					funcContent += fmt.Sprintf("// @Summary%s \n", nameArr[1])
					funcContent += fmt.Sprintf("// @Description%s \n", nameArr[1])
					// Tags
					funcContent += fmt.Sprintf("%s \n", v.Decorations().Start.All()[1])
					funcContent += fmt.Sprintf("// @Accept json \n")
					funcContent += fmt.Sprintf("// @Produce json \n")
					funcContent += fmt.Sprintf("// @Param Request body %sRequest true \"请求信息\" \n", Lcfirst(v.Names[0].String()))
					funcContent += fmt.Sprintf("// @Success 200 {object} %sResponse \n", Lcfirst(v.Names[0].String()))
					funcContent += fmt.Sprintf("// @Failure 400 {object} code.Failure \n")
					// Router
					funcContent += fmt.Sprintf("%s \n", v.Decorations().Start.All()[2])
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
