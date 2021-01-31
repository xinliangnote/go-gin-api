package trace

type Debug struct {
	Key         string      `json:"key"`          // 标示
	Value       interface{} `json:"value"`        // 值
	CostSeconds float64     `json:"cost_seconds"` // 执行时间(单位秒)
}
