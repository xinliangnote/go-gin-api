package journal

import (
	"crypto/rand"
	"encoding/hex"
	"io"
	"sync"
)

// JournalHeader http/rpc header中的名字
const JournalHeader = "Journal-ID"

var _ T = (*Journal)(nil)
var _ T = (*Dialog)(nil)

// T 约束
type T interface {
	ID() string
	t()
}

// Journal 包含一次rpc请求的全部参数和内部调用其它方接口的过程
type Journal struct {
	mux         sync.Mutex
	Identifier  string    `json:"id"`
	Request     *Request  `json:"request"`
	Response    *Response `json:"response"`
	Dialogs     []*Dialog `json:"dialogs"`
	Success     bool      `json:"success"`
	CostSeconds float64   `json:"cost_seconds"`
}

// NewJournal 创建Journal
func NewJournal(id string) *Journal {
	if id == "" {
		buf := make([]byte, 10)
		io.ReadFull(rand.Reader, buf)
		id = string(hex.EncodeToString(buf))
	}

	return &Journal{
		Identifier: id,
	}
}

// ID 唯一标识符
func (j *Journal) ID() string {
	return j.Identifier
}

// WithRequest 设置request
func (j *Journal) WithRequest(req *Request) *Journal {
	j.Request = req
	return j
}

// WithResponse 设置response
func (j *Journal) WithResponse(resp *Response) *Journal {
	j.Response = resp
	return j
}

// AppendDialog 安全的追加内部调用过程dialog
func (j *Journal) AppendDialog(dialog *Dialog) *Journal {
	if dialog == nil {
		return j
	}

	j.mux.Lock()
	defer j.mux.Unlock()

	j.Dialogs = append(j.Dialogs, dialog)
	return j
}

func (j *Journal) t() {}

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
	Header      interface{} `json:"header"`
	StatusCode  int         `json:"status_code"`
	Status      string      `json:"status"`
	Body        interface{} `json:"body"`
	CostSeconds float64     `json:"cost_seconds"`
}

// Dialog 内部调用其它方接口的会话信息;失败时会有retry操作，所以response会有多次。
type Dialog struct {
	mux         sync.Mutex
	Request     *Request    `json:"request"`
	Responses   []*Response `json:"responses"`
	Success     bool        `json:"success"`
	CostSeconds float64     `json:"cost_seconds"`
}

// ID ...
func (d *Dialog) ID() string {
	return ""
}

// AppendResponse 按转的追加response信息
func (d *Dialog) AppendResponse(resp *Response) {
	if resp == nil {
		return
	}

	d.mux.Lock()
	d.Responses = append(d.Responses, resp)
	d.mux.Unlock()
}

func (d *Dialog) t() {}
