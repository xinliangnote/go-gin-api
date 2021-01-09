package main

import (
	"fmt"

	"github.com/xinliangnote/go-gin-api/pkg/env"
)

func main() {
	fmt.Println(env.Active().Value())
}
