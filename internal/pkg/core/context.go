package core

import (
	"bytes"
	stdctx "context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/xinliangnote/go-gin-api/internal/pkg/trace"
	"github.com/xinliangnote/go-gin-api/pkg/errno"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/zap"
)

type HandlerFunc func(c Context)

type Trace = trace.T

const (
	_Alias          = "_alias_"
	_TraceName      = "_trace_"
	_LoggerName     = "_logger_"
	_BodyName       = "_body_"
	_PayloadName    = "_payload_"
	_UserID         = "_user_id_"
	_UserName       = "_user_name_"
	_AbortErrorName = "_abort_error_"
)

var contextPool = &sync.Pool{
	New: func() interface{} {
		return new(context)
	},
}

func newContext(ctx *gin.Context) Context {
	context := contextPool.Get().(*context)
	context.ctx = ctx
	return context
}

func releaseContext(ctx Context) {
	c := ctx.(*context)
	c.ctx = nil
	contextPool.Put(c)
}

var _ Context = (*context)(nil)

type Context interface {
	init()

	// ShouldBindQuery 反序列化 querystring
	// tag: `form:"xxx"` (注：不要写成 query)
	ShouldBindQuery(obj interface{}) error

	// ShouldBindPostForm 反序列化 postform (querystring会被忽略)
	// tag: `form:"xxx"`
	ShouldBindPostForm(obj interface{}) error

	// ShouldBindForm 同时反序列化 querystring 和 postform;
	// 当 querystring 和 postform 存在相同字段时，postform 优先使用。
	// tag: `form:"xxx"`
	ShouldBindForm(obj interface{}) error

	// ShouldBindJSON 反序列化 postjson
	// tag: `json:"xxx"`
	ShouldBindJSON(obj interface{}) error

	// ShouldBindURI 反序列化 path 参数(如路由路径为 /user/:name)
	// tag: `uri:"xxx"`
	ShouldBindURI(obj interface{}) error

	// Redirect 重定向
	Redirect(code int, location string)

	// Trace 获取 Trace 对象
	Trace() Trace
	setTrace(trace Trace)
	disableTrace()

	// Logger 获取 Logger 对象
	Logger() *zap.Logger
	setLogger(logger *zap.Logger)

	// Payload 正确返回
	Payload(payload errno.Error)
	getPayload() errno.Error

	// AbortWithError 错误返回
	AbortWithError(err errno.Error)
	abortError() errno.Error

	// Header 获取 Header 对象
	Header() http.Header
	// GetHeader 获取 Header
	GetHeader(key string) string
	// SetHeader 设置 Header
	SetHeader(key, value string)

	// UserID 获取 JWT 中 UserID
	UserID() int64
	setUserID(userID int64)

	// UserName 获取 JWT 中 UserName
	UserName() string
	setUserName(userName string)

	// Alias 设置路由别名 for metrics uri
	Alias() string
	setAlias(path string)

	// RawData 获取 Request.Body
	RawData() []byte
	// Method 获取 Request.Method
	Method() string
	// Host 获取 Request.Host
	Host() string
	// Path 获取 请求的路径 Request.URL.Path (不附带 querystring)
	Path() string
	// URI 获取 unescape 后的 Request.URL.RequestURI()
	URI() string
	// RequestContext 获取请求的 context (当 client 关闭后，会自动 canceled)
	RequestContext() StdContext
}

type context struct {
	ctx *gin.Context
}

type StdContext struct {
	stdctx.Context
	Trace
	*zap.Logger
}

func (c *context) init() {
	body, err := c.ctx.GetRawData()
	if err != nil {
		panic(err)
	}

	c.ctx.Set(_BodyName, body)                                   // cache body是为了trace使用
	c.ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body)) // re-construct req body
}

// ShouldBindQuery 反序列化querystring
// tag: `form:"xxx"` (注：不要写成query)
func (c *context) ShouldBindQuery(obj interface{}) error {
	return c.ctx.ShouldBindWith(obj, binding.Query)
}

// ShouldBindPostForm 反序列化 postform (querystring 会被忽略)
// tag: `form:"xxx"`
func (c *context) ShouldBindPostForm(obj interface{}) error {
	return c.ctx.ShouldBindWith(obj, binding.FormPost)
}

// ShouldBindForm 同时反序列化querystring和postform;
// 当querystring和postform存在相同字段时，postform优先使用。
// tag: `form:"xxx"`
func (c *context) ShouldBindForm(obj interface{}) error {
	return c.ctx.ShouldBindWith(obj, binding.Form)
}

// ShouldBindJSON 反序列化postjson
// tag: `json:"xxx"`
func (c *context) ShouldBindJSON(obj interface{}) error {
	return c.ctx.ShouldBindWith(obj, binding.JSON)
}

// ShouldBindURI 反序列化path参数(如路由路径为 /user/:name)
// tag: `uri:"xxx"`
func (c *context) ShouldBindURI(obj interface{}) error {
	return c.ctx.ShouldBindUri(obj)
}

// Redirect 重定向
func (c *context) Redirect(code int, location string) {
	c.ctx.Redirect(code, location)
}

func (c *context) Trace() Trace {
	t, ok := c.ctx.Get(_TraceName)
	if !ok || t == nil {
		return nil
	}

	return t.(Trace)
}

func (c *context) setTrace(trace Trace) {
	c.ctx.Set(_TraceName, trace)
}

func (c *context) disableTrace() {
	c.setTrace(nil)
}

func (c *context) Logger() *zap.Logger {
	logger, ok := c.ctx.Get(_LoggerName)
	if !ok {
		return nil
	}

	return logger.(*zap.Logger)
}

func (c *context) setLogger(logger *zap.Logger) {
	c.ctx.Set(_LoggerName, logger)
}

func (c *context) getPayload() errno.Error {
	payload, _ := c.ctx.Get(_PayloadName)
	return payload.(errno.Error)
}

func (c *context) Payload(payload errno.Error) {
	c.ctx.Set(_PayloadName, payload)
}

func (c *context) Header() http.Header {
	header := c.ctx.Request.Header

	clone := make(http.Header, len(header))
	for k, v := range header {
		value := make([]string, len(v))
		copy(value, v)

		clone[k] = value
	}
	return clone
}

func (c *context) GetHeader(key string) string {
	return c.ctx.GetHeader(key)
}

func (c *context) SetHeader(key, value string) {
	c.ctx.Header(key, value)
}

func (c *context) UserID() int64 {
	val, ok := c.ctx.Get(_UserID)
	if !ok {
		return 0
	}

	return val.(int64)
}

func (c *context) setUserID(userID int64) {
	c.ctx.Set(_UserID, userID)
}

func (c *context) UserName() string {
	val, ok := c.ctx.Get(_UserName)
	if !ok {
		return ""
	}

	return val.(string)
}

func (c *context) setUserName(userName string) {
	c.ctx.Set(_UserName, userName)
}

func (c *context) AbortWithError(err errno.Error) {
	if err != nil {
		httpCode := err.GetHttpCode()
		if httpCode == 0 {
			httpCode = http.StatusInternalServerError
		}

		c.ctx.AbortWithStatus(httpCode)
		c.ctx.Set(_AbortErrorName, err)
	}
}

func (c *context) abortError() errno.Error {
	err, _ := c.ctx.Get(_AbortErrorName)
	return err.(errno.Error)
}

func (c *context) Alias() string {
	path, ok := c.ctx.Get(_Alias)
	if !ok {
		return ""
	}

	return path.(string)
}

func (c *context) setAlias(path string) {
	if path = strings.TrimSpace(path); path != "" {
		c.ctx.Set(_Alias, path)
	}
}

func (c *context) RawData() []byte {
	body, ok := c.ctx.Get(_BodyName)
	if !ok {
		return nil
	}

	return body.([]byte)
}

// Method 请求的method
func (c *context) Method() string {
	return c.ctx.Request.Method
}

// Host 请求的host
func (c *context) Host() string {
	return c.ctx.Request.Host
}

// Path 请求的路径(不附带querystring)
func (c *context) Path() string {
	return c.ctx.Request.URL.Path
}

// URI unescape后的uri
func (c *context) URI() string {
	uri, _ := url.QueryUnescape(c.ctx.Request.URL.RequestURI())
	return uri
}

// RequestContext 获取请求的 context (当client关闭后，会自动canceled)
func (c *context) RequestContext() StdContext {
	return StdContext{
		c.ctx.Request.Context(),
		c.Trace(),
		c.Logger(),
	}
}
