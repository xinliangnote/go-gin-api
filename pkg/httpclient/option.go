package httpclient

import (
	"sync"
	"time"

	"github.com/xinliangnote/go-gin-api/pkg/trace"

	"go.uber.org/zap"
)

var (
	cache = &sync.Pool{
		New: func() interface{} {
			return &option{
				header: make(map[string][]string),
			}
		},
	}
)

// Mock 定义接口Mock数据
type Mock func() (body []byte)

// Option 自定义设置http请求
type Option func(*option)

type option struct {
	ttl         time.Duration
	header      map[string][]string
	trace       *trace.Trace
	dialog      *trace.Dialog
	logger      *zap.Logger
	retryTimes  int
	retryDelay  time.Duration
	retryVerify RetryVerify
	alarmTitle  string
	alarmObject AlarmObject
	alarmVerify AlarmVerify
	mock        Mock
}

func (o *option) reset() {
	o.ttl = 0
	o.header = make(map[string][]string)
	o.trace = nil
	o.dialog = nil
	o.logger = nil
	o.retryTimes = 0
	o.retryDelay = 0
	o.retryVerify = nil
	o.alarmTitle = ""
	o.alarmObject = nil
	o.alarmVerify = nil
	o.mock = nil
}

func getOption() *option {
	return cache.Get().(*option)
}

func releaseOption(opt *option) {
	opt.reset()
	cache.Put(opt)
}

// WithTTL 本次http请求最长执行时间
func WithTTL(ttl time.Duration) Option {
	return func(opt *option) {
		opt.ttl = ttl
	}
}

// WithHeader 设置http header，可以调用多次设置多对key-value
func WithHeader(key, value string) Option {
	return func(opt *option) {
		opt.header[key] = []string{value}
	}
}

// WithTrace 设置trace信息
func WithTrace(t trace.T) Option {
	return func(opt *option) {
		if t != nil {
			opt.trace = t.(*trace.Trace)
			opt.dialog = new(trace.Dialog)
		}
	}
}

// WithLogger 设置logger以便打印关键日志
func WithLogger(logger *zap.Logger) Option {
	return func(opt *option) {
		opt.logger = logger
	}
}

// WithMock 设置 mock 数据
func WithMock(m Mock) Option {
	return func(opt *option) {
		opt.mock = m
	}
}

// WithOnFailedAlarm 设置告警通知
func WithOnFailedAlarm(alarmTitle string, alarmObject AlarmObject, alarmVerify AlarmVerify) Option {
	return func(opt *option) {
		opt.alarmTitle = alarmTitle
		opt.alarmObject = alarmObject
		opt.alarmVerify = alarmVerify
	}
}

// WithOnFailedRetry 设置失败重试
func WithOnFailedRetry(retryTimes int, retryDelay time.Duration, retryVerify RetryVerify) Option {
	return func(opt *option) {
		opt.retryTimes = retryTimes
		opt.retryDelay = retryDelay
		opt.retryVerify = retryVerify
	}
}
