package request

import (
	"crypto/tls"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"go-gin-api/app/config"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func HttpGet(url string, c *gin.Context) (string, error) {

	tr := &http.Transport{
		TLSClientConfig : &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Timeout   : time.Second * 5, //默认5秒超时时间
		Transport : tr,
	}

	req, err := http.NewRequest("GET", url,nil)
	if err != nil {
		return "", err
	}

	if config.JaegerOpen == 1 {

		tracer, _            := c.Get("Tracer")
		parentSpanContext, _ := c.Get("ParentSpanContext")

		span := opentracing.StartSpan(
			"call Http Get",
			opentracing.ChildOf(parentSpanContext.(opentracing.SpanContext)),
			opentracing.Tag{Key: string(ext.Component), Value: "HTTP"},
			ext.SpanKindRPCClient,
		)

		span.Finish()

		injectErr := tracer.(opentracing.Tracer).Inject(span.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(req.Header))
		if injectErr != nil {
			log.Fatalf("%s: Couldn't inject headers", err)
		}
	}

	resp ,err :=  client.Do(req)
	if err != nil {
		return "", err
	}
	content, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		return "", err
	}
	return string(content), err
}
