package main

import (
	"flag"
	"log"
	"os"
	"strings"

	"github.com/xinliangnote/go-gin-api/cmd/gormgen/pkg"
)

var (
	input   string
	structs []string
)

func init() {
	flagStructs := flag.String("structs", "", "[Required] The name of schema structs to generate structs for, comma seperated\n")
	flagInput := flag.String("input", "", "[Required] The name of the input file dir\n")
	flag.Parse()

	if *flagStructs == "" || *flagInput == "" {
		flag.Usage()
		os.Exit(1)
	}

	structs = strings.Split(*flagStructs, ",")
	input = *flagInput
}

func main() {
	gen := pkg.NewGenerator(input)
	p := pkg.NewParser(input)
	if err := gen.ParserAST(p, structs).Generate().Format().Flush(); err != nil {
		log.Fatalln(err)
	}
}
