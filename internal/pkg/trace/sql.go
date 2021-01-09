package trace

type SQL struct {
	Time     string  `json:"time"`          // 时间，格式：2006-01-02 15:04:05
	Src      string  `json:"src"`           // 文件地址和行号
	Duration float64 `json:"duration"`      // 执行时长，单位：秒
	SQL      string  `json:"sql"`           // SQL 语句
	Rows     int64   `json:"rows_affected"` // 影响行数
}
