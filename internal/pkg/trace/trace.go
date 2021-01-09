package trace

import (
	"crypto/rand"
	"encoding/hex"
	"io"
	"sync"
)

const Header = "TRACE-ID"

var _ T = (*Trace)(nil)

type T interface {
	i()
	ID() string
	WithRequest(req *Request) *Trace
	WithResponse(resp *Response) *Trace
	AppendDialog(dialog *Dialog) *Trace
	AppendSQL(sql *SQL) *Trace
}

// Trace 记录的参数
type Trace struct {
	mux                sync.Mutex
	Identifier         string    `json:"trace_id"`
	Request            *Request  `json:"request"`
	Response           *Response `json:"response"`
	ThirdPartyRequests []*Dialog `json:"third_party_requests"`
	Debugs             []*Debug  `json:"debugs"`
	SQLs               []*SQL    `json:"sqls"`
	Success            bool      `json:"success"`
	CostSeconds        float64   `json:"cost_seconds"`
}

// Request 请求信息
type Request struct {
	TTL        string      `json:"ttl"`
	Method     string      `json:"method"`
	DecodedURL string      `json:"decoded_url"`
	Header     interface{} `json:"header"`
	Body       interface{} `json:"body"`
}

// Response 响应信息
type Response struct {
	Header          interface{} `json:"header"`
	Body            interface{} `json:"body"`
	BusinessCode    int         `json:"business_code,omitempty"`
	BusinessCodeMsg string      `json:"business_code_msg,omitempty"`
	HttpCode        int         `json:"http_code"`
	HttpCodeMsg     string      `json:"http_code_msg"`
	CostSeconds     float64     `json:"cost_seconds"`
}

func New(id string) *Trace {
	if id == "" {
		buf := make([]byte, 10)
		io.ReadFull(rand.Reader, buf)
		id = string(hex.EncodeToString(buf))
	}

	return &Trace{
		Identifier: id,
	}
}

func (t *Trace) i() {}

// ID 唯一标识符
func (t *Trace) ID() string {
	return t.Identifier
}

// WithRequest 设置request
func (t *Trace) WithRequest(req *Request) *Trace {
	t.Request = req
	return t
}

// WithResponse 设置response
func (t *Trace) WithResponse(resp *Response) *Trace {
	t.Response = resp
	return t
}

// AppendDialog 安全的追加内部调用过程dialog
func (t *Trace) AppendDialog(dialog *Dialog) *Trace {
	if dialog == nil {
		return t
	}

	t.mux.Lock()
	defer t.mux.Unlock()

	t.ThirdPartyRequests = append(t.ThirdPartyRequests, dialog)
	return t
}

// AppendDebug 追加 debug
func (t *Trace) AppendDebug(debug *Debug) *Trace {
	if debug == nil {
		return t
	}

	t.mux.Lock()
	defer t.mux.Unlock()

	t.Debugs = append(t.Debugs, debug)
	return t
}

// AppendSQL 追加 SQL
func (t *Trace) AppendSQL(sql *SQL) *Trace {
	if sql == nil {
		return t
	}

	t.mux.Lock()
	defer t.mux.Unlock()

	t.SQLs = append(t.SQLs, sql)
	return t
}
