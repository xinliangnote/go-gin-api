package jaeger

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"go-gin-api/app/config"
	"go-gin-api/app/util/jaeger_trace"
	"io"
)

var Tracer opentracing.Tracer
var Closer io.Closer
var Error  error

var ParentSpan    opentracing.Span

func SetUp() gin.HandlerFunc {

	return func(c *gin.Context) {
		if config.JaegerOpen == 1 {
			Tracer, Closer, Error = jaeger_trace.NewJaegerTracer(config.AppName, config.JaegerHostPort)
			defer Closer.Close()

			spCtx, err := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
			if err != nil {
				ParentSpan = Tracer.StartSpan(c.Request.URL.Path)
				defer ParentSpan.Finish()
			} else {
				ParentSpan = opentracing.StartSpan(
					c.Request.URL.Path,
					opentracing.ChildOf(spCtx),
					opentracing.Tag{Key: string(ext.Component), Value: "HTTP"},
					ext.SpanKindRPCServer,
				)
				defer ParentSpan.Finish()
			}
		}
		c.Next()
	}
}
