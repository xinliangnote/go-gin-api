package httpclient

import (
	"time"

	"github.com/xinliangnote/go-gin-api/internal/pkg/journal"

	"go.uber.org/zap"
)

// Journal 记录内部流转信息
type Journal = journal.T

// Option 自定义设置http请求
type Option func(*option)

type option struct {
	TTL        time.Duration
	Header     map[string]string
	Journal    *journal.Journal
	Dialog     *journal.Dialog
	Logger     *zap.Logger
	RetryTimes int
	RetryDelay time.Duration
}

func newOption() *option {
	return &option{
		Header: make(map[string]string),
	}
}

// WithTTL 本次http请求最长执行时间
func WithTTL(ttl time.Duration) Option {
	return func(opt *option) {
		opt.TTL = ttl
	}
}

// WithHeader 设置http header，可以调用多次设置多对key-value
func WithHeader(key, value string) Option {
	return func(opt *option) {
		opt.Header[key] = value
	}
}

// WithJournal 设置Journal以便记录内部流转信息
func WithJournal(j Journal) Option {
	return func(opt *option) {
		if j != nil {
			opt.Journal = j.(*journal.Journal)
			opt.Dialog = new(journal.Dialog)
		}
	}
}

// WithLogger 设置logger以便打印关键日志
func WithLogger(logger *zap.Logger) Option {
	return func(opt *option) {
		opt.Logger = logger
	}
}

// WithRetryTimes 如果请求失败，最多重试N次
func WithRetryTimes(retryTimes int) Option {
	return func(opt *option) {
		opt.RetryTimes = retryTimes
	}
}

// WithRetryDelay 在重试前，延迟等待一会
func WithRetryDelay(retryDelay time.Duration) Option {
	return func(opt *option) {
		opt.RetryDelay = retryDelay
	}
}
