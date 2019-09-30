package jaeger

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"go-gin-api/app/config"
	"go-gin-api/app/util/jaeger_trace"
)

func SetUp() gin.HandlerFunc {

	return func(c *gin.Context) {
		if config.JaegerOpen == 1 {

			var parentSpan opentracing.Span

			tracer, closer := jaeger_trace.NewJaegerTracer(config.AppName, config.JaegerHostPort)
			defer closer.Close()

			spCtx, err := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
			if err != nil {
				parentSpan = tracer.StartSpan(c.Request.URL.Path)
				defer parentSpan.Finish()
			} else {
				parentSpan = opentracing.StartSpan(
					c.Request.URL.Path,
					opentracing.ChildOf(spCtx),
					opentracing.Tag{Key: string(ext.Component), Value: "HTTP"},
					ext.SpanKindRPCServer,
				)
				defer parentSpan.Finish()
			}
			c.Set("Tracer", tracer)
			c.Set("ParentSpanContext", parentSpan.Context())
		}
		c.Next()
	}
}
