package debugs

import (
	"fmt"
	"time"

	"github.com/xinliangnote/go-gin-api/pkg/trace"
)

type Option func(*option)

type Trace = trace.T

type option struct {
	Trace *trace.Trace
	Debug *trace.Debug
}

func newOption() *option {
	return &option{}
}

func Println(key string, value interface{}, options ...Option) {
	ts := time.Now()
	opt := newOption()
	defer func() {
		if opt.Trace != nil {
			opt.Debug.Key = key
			opt.Debug.Value = value
			opt.Debug.CostSeconds = time.Since(ts).Seconds()
			opt.Trace.AppendDebug(opt.Debug)
		}
	}()

	for _, f := range options {
		f(opt)
	}

	fmt.Println(fmt.Sprintf("KEY: %s | VALUE: %v", key, value))
}

// WithTrace 设置trace信息
func WithTrace(t Trace) Option {
	return func(opt *option) {
		if t != nil {
			opt.Trace = t.(*trace.Trace)
			opt.Debug = new(trace.Debug)
		}
	}
}
