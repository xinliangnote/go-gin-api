package trace

import "sync"

var _ D = (*Dialog)(nil)

type D interface {
	i()
	AppendResponse(resp *Response)
}

// Dialog 内部调用其它方接口的会话信息；失败时会有retry操作，所以 response 会有多次。
type Dialog struct {
	mux         sync.Mutex
	Request     *Request    `json:"request"`      // 请求信息
	Responses   []*Response `json:"responses"`    // 返回信息
	Success     bool        `json:"success"`      // 是否成功，true 或 false
	CostSeconds float64     `json:"cost_seconds"` // 执行时长(单位秒)
}

func (d *Dialog) i() {}

// AppendResponse 按转的追加response信息
func (d *Dialog) AppendResponse(resp *Response) {
	if resp == nil {
		return
	}

	d.mux.Lock()
	d.Responses = append(d.Responses, resp)
	d.mux.Unlock()
}
