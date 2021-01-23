package trace

import "time"

type Redis struct {
	Timestamp string        `json:"timestamp"` // 时间，格式：2006-01-02 15:04:05
	Handle    string        `json:"handle"`    // 操作，function
	Key       string        `json:"key"`
	Value     string        `json:"value,omitempty"`
	TTL       time.Duration `json:"ttl,omitempty"`
}
