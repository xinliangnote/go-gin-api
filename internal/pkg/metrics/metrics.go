package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/cast"
)

// requestsCounter 定义计数器（Counter）
var requestsCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "api_requests_total",
	},
	[]string{"method", "path"},
)

// httpDurationsHistogram 定义累积直方图（Histogram）
var httpDurationsHistogram = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "api_http_durations_histogram_seconds",
		Buckets: []float64{0.01, 0.02, 0.03},
	},
	[]string{"method", "path", "success", "http_code", "business_code", "cost_seconds", "trace_id"},
)

func init() {
	prometheus.MustRegister(requestsCounter, httpDurationsHistogram)
}

// RecordMetrics 记录指标
func RecordMetrics(method, uri string, success bool, httpCode, businessCode int, costSeconds float64, traceId string) {
	httpDurationsHistogram.With(prometheus.Labels{
		"method":        method,
		"path":          uri,
		"success":       cast.ToString(success),
		"http_code":     cast.ToString(httpCode),
		"business_code": cast.ToString(businessCode),
		"cost_seconds":  cast.ToString(costSeconds),
		"trace_id":      traceId,
	}).Observe(costSeconds)

	requestsCounter.With(prometheus.Labels{
		"method": method,
		"path":   uri,
	}).Add(1)
}
