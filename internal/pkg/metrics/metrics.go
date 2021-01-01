package metrics

// RecordMetrics 记录指标
func RecordMetrics(method, uri string, success bool, httpCode, businessCode int, costSeconds float64) {
	//fmt.Printf(">>>>>>>Metrics\nmethod:%s\nuri:%s\nsuccess:%t\nhttp code:%d\nbusiness code:%d\ncost seconds:%.9f\n<<<<<<<\n", method, uri, success, httpCode, businessCode, costSeconds)
}
