package pkg

import (
	"bytes"
	"errors"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"strings"

	"github.com/jinzhu/gorm"
)

// fieldConfig
type fieldConfig struct {
	FieldName  string
	ColumnName string
	FieldType  string
	HumpName   string
}

// structConfig
type structConfig struct {
	config
	StructName   string
	OnlyFields   []fieldConfig
	OptionFields []fieldConfig
}

type ImportPkg struct {
	Pkg string
}

type structHelpers struct {
	Titelize func(string) string
}

type config struct {
	PkgName          string
	Helpers          structHelpers
	QueryBuilderName string
}

// The Generator is the one responsible for generating the code, adding the imports, formating, and writing it to the file.
type Generator struct {
	buf           map[string]*bytes.Buffer
	inputFile     string
	config        config
	structConfigs []structConfig
}

// NewGenerator function creates an instance of the generator given the name of the output file as an argument.
func NewGenerator(outputFile string) *Generator {
	return &Generator{
		buf:       map[string]*bytes.Buffer{},
		inputFile: outputFile,
	}
}

// ParserAST parse by go file
func (g *Generator) ParserAST(p *Parser, structs []string) (ret *Generator) {
	for _, v := range structs {
		g.buf[gorm.ToDBName(v)] = new(bytes.Buffer)
	}
	g.structConfigs = p.Parse()
	g.config.PkgName = p.pkg.Name
	g.config.Helpers = structHelpers{
		Titelize: strings.Title,
	}
	g.config.QueryBuilderName = SQLColumnToHumpStyle(p.pkg.Name) + "QueryBuilder"
	return g
}

func (g *Generator) checkConfig() (err error) {
	if len(g.config.PkgName) == 0 {
		err = errors.New("package name dose'n set")
		return
	}
	for i := 0; i < len(g.structConfigs); i++ {
		g.structConfigs[i].config = g.config
	}
	return
}

// Generate executes the template and store it in an internal buffer.
func (g *Generator) Generate() *Generator {
	if err := g.checkConfig(); err != nil {
		panic(err)
	}

	for _, v := range g.structConfigs {
		if _, ok := g.buf[gorm.ToDBName(v.StructName)]; !ok {
			continue
		}
		if err := outputTemplate.Execute(g.buf[gorm.ToDBName(v.StructName)], v); err != nil {
			panic(err)
		}
	}

	return g
}

// Format function formats the output of the generation.
func (g *Generator) Format() *Generator {
	for k := range g.buf {
		formattedOutput, err := format.Source(g.buf[k].Bytes())
		if err != nil {
			panic(err)
		}
		g.buf[k] = bytes.NewBuffer(formattedOutput)
	}
	return g
}

// Flush function writes the output to the output file.
func (g *Generator) Flush() error {
	for k := range g.buf {
		filename := g.inputFile + "/gen_" + strings.ToLower(k) + ".go"
		if err := ioutil.WriteFile(filename, g.buf[k].Bytes(), 0777); err != nil {
			log.Fatalln(err)
		}
		fmt.Println("  └── file : ", fmt.Sprintf("%s_repo/gen_%s.go", strings.ToLower(k), strings.ToLower(k)))
	}
	return nil
}
