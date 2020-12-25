package httpclient

import (
	"context"
	"encoding/json"
	"net/http"
	httpURL "net/url"
	"time"

	"github.com/xinliangnote/go-gin-api/internal/pkg/journal"

	"github.com/pkg/errors"
)

const (
	// DefaultTTL 一次http请求最长执行1分钟
	DefaultTTL = time.Minute
	// DefaultRetryTimes 如果请求失败，最多重试3次
	DefaultRetryTimes = 3
	// DefaultRetryDelay 在重试前，延迟等待100毫秒
	DefaultRetryDelay = time.Millisecond * 100
)

// TODO retry的code不一定正确，缺失或者多余待实际使用中修改。
func shouldRetry(ctx context.Context, httpCode int) bool {
	select {
	case <-ctx.Done():
		return false
	default:
	}

	switch httpCode {
	case
		_StatusDoReqErr,    // customize
		_StatusReadRespErr, // customize

		http.StatusRequestTimeout,
		http.StatusLocked,
		http.StatusTooEarly,
		http.StatusTooManyRequests,

		http.StatusServiceUnavailable,
		http.StatusGatewayTimeout:

		return true

	default:
		return false
	}
}

// Get get 请求
func Get(url string, form httpURL.Values, options ...Option) (body []byte, err error) {
	return withoutBody(http.MethodGet, url, form, options...)
}

// Delete delete 请求
func Delete(url string, form httpURL.Values, options ...Option) (body []byte, err error) {
	return withoutBody(http.MethodDelete, url, form, options...)
}

func withoutBody(method, url string, form httpURL.Values, options ...Option) (body []byte, err error) {
	if url == "" {
		return nil, errors.New("url required")
	}

	if len(form) > 0 {
		if url, err = addFormValuesIntoURL(url, form); err != nil {
			return
		}
	}

	ts := time.Now()

	opt := newOption()
	defer func() {
		if opt.Journal != nil {
			opt.Dialog.Success = err == nil
			opt.Dialog.CostSeconds = time.Since(ts).Seconds()
			opt.Journal.AppendDialog(opt.Dialog)
		}
	}()

	for _, f := range options {
		f(opt)
	}
	opt.Header["Content-Type"] = "application/x-www-form-urlencoded; charset=utf-8"
	if opt.Journal != nil {
		opt.Header[journal.JournalHeader] = opt.Journal.ID()
	}

	ttl := opt.TTL
	if ttl <= 0 {
		ttl = DefaultTTL
	}

	ctx, cancel := context.WithTimeout(context.Background(), ttl)
	defer cancel()

	if opt.Dialog != nil {
		decodedURL, _ := httpURL.QueryUnescape(url)
		opt.Dialog.Request = &journal.Request{
			TTL:        ttl.String(),
			Method:     method,
			DecodedURL: decodedURL,
			Header:     opt.Header,
		}
	}

	retryTimes := opt.RetryTimes
	if retryTimes <= 0 {
		retryTimes = DefaultRetryTimes
	}

	retryDelay := opt.RetryDelay
	if retryDelay <= 0 {
		retryDelay = DefaultRetryDelay
	}

	var httpCode int
	for k := 0; k < retryTimes; k++ {
		body, httpCode, err = doHTTP(ctx, method, url, nil, opt)
		if shouldRetry(ctx, httpCode) {
			time.Sleep(retryDelay)
			continue
		}

		return
	}
	return
}

// PostForm post form 请求
func PostForm(url string, form httpURL.Values, options ...Option) (body []byte, err error) {
	return withFormBody(http.MethodPost, url, form, options...)
}

// PostJSON post json 请求
func PostJSON(url string, raw json.RawMessage, options ...Option) (body []byte, err error) {
	return withJSONBody(http.MethodPost, url, raw, options...)
}

// PutForm put form 请求
func PutForm(url string, form httpURL.Values, options ...Option) (body []byte, err error) {
	return withFormBody(http.MethodPut, url, form, options...)
}

// PutJSON put json 请求
func PutJSON(url string, raw json.RawMessage, options ...Option) (body []byte, err error) {
	return withJSONBody(http.MethodPut, url, raw, options...)
}

// PatchFrom patch form 请求
func PatchFrom(url string, form httpURL.Values, options ...Option) (body []byte, err error) {
	return withFormBody(http.MethodPatch, url, form, options...)
}

// PatchJSON patch json 请求
func PatchJSON(url string, raw json.RawMessage, options ...Option) (body []byte, err error) {
	return withJSONBody(http.MethodPatch, url, raw, options...)
}

func withFormBody(method, url string, form httpURL.Values, options ...Option) (body []byte, err error) {
	if url == "" {
		return nil, errors.New("url required")
	}
	if len(form) == 0 {
		return nil, errors.New("form required")
	}

	ts := time.Now()

	opt := newOption()
	defer func() {
		if opt.Journal != nil {
			opt.Dialog.Success = err == nil
			opt.Dialog.CostSeconds = time.Since(ts).Seconds()
			opt.Journal.AppendDialog(opt.Dialog)
		}
	}()

	for _, f := range options {
		f(opt)
	}
	opt.Header["Content-Type"] = "application/x-www-form-urlencoded; charset=utf-8"
	if opt.Journal != nil {
		opt.Header[journal.JournalHeader] = opt.Journal.ID()
	}

	ttl := opt.TTL
	if ttl <= 0 {
		ttl = DefaultTTL
	}

	ctx, cancel := context.WithTimeout(context.Background(), ttl)
	defer cancel()

	formValue := form.Encode()
	if opt.Dialog != nil {
		decodedURL, _ := httpURL.QueryUnescape(url)
		opt.Dialog.Request = &journal.Request{
			TTL:        ttl.String(),
			Method:     method,
			DecodedURL: decodedURL,
			Header:     opt.Header,
			Body:       formValue,
		}
	}

	retryTimes := opt.RetryTimes
	if retryTimes <= 0 {
		retryTimes = DefaultRetryTimes
	}

	retryDelay := opt.RetryDelay
	if retryDelay <= 0 {
		retryDelay = DefaultRetryDelay
	}

	var httpCode int
	for k := 0; k < retryTimes; k++ {
		body, httpCode, err = doHTTP(ctx, method, url, []byte(formValue), opt)
		if shouldRetry(ctx, httpCode) {
			time.Sleep(retryDelay)
			continue
		}

		return
	}
	return
}

func withJSONBody(method, url string, raw json.RawMessage, options ...Option) (body []byte, err error) {
	if url == "" {
		return nil, errors.New("url required")
	}
	if len(raw) == 0 {
		return nil, errors.New("raw required")
	}

	ts := time.Now()

	opt := newOption()
	defer func() {
		if opt.Journal != nil {
			opt.Dialog.Success = err == nil
			opt.Dialog.CostSeconds = time.Since(ts).Seconds()
			opt.Journal.AppendDialog(opt.Dialog)
		}
	}()

	for _, f := range options {
		f(opt)
	}
	opt.Header["Content-Type"] = "application/json; charset=utf-8"
	if opt.Journal != nil {
		opt.Header[journal.JournalHeader] = opt.Journal.ID()
	}

	ttl := opt.TTL
	if ttl <= 0 {
		ttl = DefaultTTL
	}

	ctx, cancel := context.WithTimeout(context.Background(), ttl)
	defer cancel()

	if opt.Dialog != nil {
		decodedURL, _ := httpURL.QueryUnescape(url)
		opt.Dialog.Request = &journal.Request{
			TTL:        ttl.String(),
			Method:     method,
			DecodedURL: decodedURL,
			Header:     opt.Header,
			Body:       string(raw), // TODO unsafe
		}
	}

	retryTimes := opt.RetryTimes
	if retryTimes <= 0 {
		retryTimes = DefaultRetryTimes
	}

	retryDelay := opt.RetryDelay
	if retryDelay <= 0 {
		retryDelay = DefaultRetryDelay
	}

	var httpCode int
	for k := 0; k < retryTimes; k++ {
		body, httpCode, err = doHTTP(ctx, method, url, raw, opt)
		if shouldRetry(ctx, httpCode) {
			time.Sleep(retryDelay)
			continue
		}

		return
	}
	return
}
