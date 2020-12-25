package errno

import (
	"encoding/json"
)

var _ Error = (*err)(nil)

type Error interface {
	// i 为了避免被其他包实现
	i()
	// WithData 设置成功时返回的数据
	WithData(data interface{}) Error
	// WithID 设置当前请求的唯一ID
	WithID(id string) Error
	// ToString 返回 JSON 格式的错误详情
	ToString() string
}

type err struct {
	Code int         `json:"code"`         // 业务编码
	Msg  string      `json:"msg"`          // 错误描述
	Data interface{} `json:"data"`         // 成功时返回的数据
	ID   string      `json:"id,omitempty"` // 当前请求的唯一ID，便于问题定位，忽略也可以
}

func NewError(code int, msg string) Error {
	return &err{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}

func (e *err) i() {}

func (e *err) WithData(data interface{}) Error {
	e.Data = data
	return e
}

func (e *err) WithID(id string) Error {
	e.ID = id
	return e
}

// ToString 返回 JSON 格式的错误详情
func (e *err) ToString() string {
	err := &struct {
		Code int         `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
		ID   string      `json:"id,omitempty"`
	}{
		Code: e.Code,
		Msg:  e.Msg,
		Data: e.Data,
		ID:   e.ID,
	}

	raw, _ := json.Marshal(err)
	return string(raw)
}
