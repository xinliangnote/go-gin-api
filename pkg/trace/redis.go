package trace

type Redis struct {
	Timestamp   string  `json:"timestamp"`       // 时间，格式：2006-01-02 15:04:05
	Handle      string  `json:"handle"`          // 操作，SET/GET 等
	Key         string  `json:"key"`             // Key
	Value       string  `json:"value,omitempty"` // Value
	TTL         float64 `json:"ttl,omitempty"`   // 超时时长(单位分)
	CostSeconds float64 `json:"cost_seconds"`    // 执行时间(单位秒)
}
