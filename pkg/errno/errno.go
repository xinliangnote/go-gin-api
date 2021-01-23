package errno

import (
	"encoding/json"

	"github.com/pkg/errors"
)

var _ Error = (*err)(nil)

type Error interface {
	// i 为了避免被其他包实现
	i()
	// WithData 设置成功时返回的数据
	WithData(data interface{}) Error
	// WithID 设置当前请求的唯一ID
	WithID(id string) Error
	// WithErr 设置错误信息
	WithErr(err error) Error
	// GetBusinessCode 获取 Business Code
	GetBusinessCode() int
	// GetHttpCode 获取 HTTP Code
	GetHttpCode() int
	// GetMsg 获取 Msg
	GetMsg() string
	// GetErr 获取错误信息
	GetErr() error
	// ToString 返回 JSON 格式的错误详情
	ToString() string
}

type err struct {
	HttpCode     int         `json:"-"`            // HTTP Code
	BusinessCode int         `json:"code"`         // Business Code
	Msg          string      `json:"msg"`          // 描述信息
	Data         interface{} `json:"data"`         // 接口数据
	Err          error       `json:"-"`            // 错误信息
	ID           string      `json:"id,omitempty"` // 当前请求的唯一ID，便于问题定位，忽略也可以
}

func NewError(httpCode, businessCode int, msg string) Error {
	return &err{
		HttpCode:     httpCode,
		BusinessCode: businessCode,
		Msg:          msg,
		Data:         nil,
		Err:          nil,
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

func (e *err) WithErr(err error) Error {
	e.Err = errors.WithStack(err)
	return e
}

func (e *err) GetHttpCode() int {
	return e.HttpCode
}

func (e *err) GetBusinessCode() int {
	return e.BusinessCode
}

func (e *err) GetMsg() string {
	return e.Msg
}

func (e *err) GetErr() error {
	return e.Err
}

// ToString 返回 JSON 格式的错误详情
func (e *err) ToString() string {
	err := &struct {
		HttpCode     int         `json:"http_code"`
		BusinessCode int         `json:"business_code"`
		Msg          string      `json:"msg"`
		Data         interface{} `json:"data"`
		ID           string      `json:"id,omitempty"`
	}{
		HttpCode:     e.HttpCode,
		BusinessCode: e.BusinessCode,
		Msg:          e.Msg,
		Data:         e.Data,
		ID:           e.ID,
	}

	raw, _ := json.Marshal(err)
	return string(raw)
}
