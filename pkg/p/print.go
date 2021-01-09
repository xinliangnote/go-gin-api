package p

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/trace"

	"github.com/davecgh/go-spew/spew"
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

func Print(key string, value interface{}, options ...Option) {
	opt := newOption()
	defer func() {
		if opt.Trace != nil {
			opt.Debug.Key = key
			opt.Debug.Value = value
			opt.Trace.AppendDebug(opt.Debug)
		}
	}()

	for _, f := range options {
		f(opt)
	}

	spew.Dump(key, value)
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
