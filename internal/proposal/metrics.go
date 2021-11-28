package proposal

import (
	"encoding/json"
)

// MetricsMessage 指标信息
type MetricsMessage struct {
	ProjectName  string  `json:"project_name"`  // 项目名，用于区分不同项目告警信息
	Env          string  `json:"env"`           // 运行环境
	TraceID      string  `json:"trace_id"`      // 唯一ID，用于追踪关联
	HOST         string  `json:"host"`          // 请求 HOST
	Path         string  `json:"path"`          // 请求 Path
	Method       string  `json:"method"`        // 请求 Method
	HTTPCode     int     `json:"http_code"`     // HTTP 状态码
	BusinessCode int     `json:"business_code"` // 业务码
	CostSeconds  float64 `json:"cost_seconds"`  // 耗时，单位：秒
	IsSuccess    bool    `json:"is_success"`    // 状态，是否成功
}

// Marshal 序列化到JSON
func (m *MetricsMessage) Marshal() (jsonRaw []byte) {
	jsonRaw, _ = json.Marshal(m)
	return
}

// RecordHandler 指标的记录句柄
type RecordHandler func(msg *MetricsMessage)
