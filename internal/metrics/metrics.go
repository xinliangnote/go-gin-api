package metrics

import (
	"github.com/xinliangnote/go-gin-api/internal/proposal"

	"go.uber.org/zap"
)

// RecordHandler 指标处理
func RecordHandler(logger *zap.Logger) func(msg *proposal.MetricsMessage) {
	if logger == nil {
		panic("logger required")
	}

	return func(msg *proposal.MetricsMessage) {
		RecordMetrics(
			msg.Method,
			msg.Path,
			msg.IsSuccess,
			msg.HTTPCode,
			msg.BusinessCode,
			msg.CostSeconds,
			msg.TraceID,
		)
	}
}
