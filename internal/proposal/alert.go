package proposal

import (
	"encoding/json"
	"time"
)

// AlertMessage 告警信息
type AlertMessage struct {
	ProjectName  string      `json:"project_name"`  // 项目名，用于区分不同项目告警信息
	Env          string      `json:"env"`           // 运行环境
	TraceID      string      `json:"trace_id"`      // 唯一ID，用于追踪关联
	HOST         string      `json:"host"`          // 请求 HOST
	URI          string      `json:"uri"`           // 请求 URI
	Method       string      `json:"method"`        // 请求 Method
	ErrorMessage interface{} `json:"error_message"` // 错误信息
	ErrorStack   string      `json:"error_stack"`   // 堆栈信息
	Timestamp    time.Time   `json:"timestamp"`     // 时间戳
}

// Marshal 序列化到JSON
func (a *AlertMessage) Marshal() (jsonRaw []byte) {
	jsonRaw, _ = json.Marshal(a)
	return
}

// NotifyHandler 告警的发送句柄
type NotifyHandler func(msg *AlertMessage)
