package trace

import (
	"google.golang.org/grpc/metadata"
)

type Grpc struct {
	Timestamp   string                 `json:"timestamp"`             // 时间，格式：2006-01-02 15:04:05
	Addr        string                 `json:"addr"`                  // 地址
	Method      string                 `json:"method"`                // 操作方法
	Meta        metadata.MD            `json:"meta"`                  // Mate
	Request     map[string]interface{} `json:"request"`               // 请求信息
	Response    map[string]interface{} `json:"response"`              // 返回信息
	CostSeconds float64                `json:"cost_seconds"`          // 执行时间(单位秒)
	Code        string                 `json:"err_code,omitempty"`    // 错误码
	Message     string                 `json:"err_message,omitempty"` // 错误信息
}
