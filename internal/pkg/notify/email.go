package notify

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/pkg/color"
)

// Email
func Email(ctx core.Context, err interface{}, stackInfo string) {

	buf := bytes.NewBuffer(nil)
	buf.WriteString(fmt.Sprintf("URI: %s %s%s <br/>", ctx.Method(), ctx.Host(), ctx.URI()))
	buf.WriteString(fmt.Sprintf("JournalID: %s <br/>", ctx.Journal().ID()))
	buf.WriteString(fmt.Sprintf("ErrorInfo: %+v <br/>", err))
	buf.WriteString(fmt.Sprintf("StackInfo: <br/>"))

	for _, v := range strings.Split(stackInfo, "\n") {
		buf.WriteString(v)
		buf.WriteString(" <br/>")
	}

	content := buf.String()

	fmt.Println(color.Red(content))
}
