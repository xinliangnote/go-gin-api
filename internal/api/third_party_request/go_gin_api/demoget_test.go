package go_gin_api

import (
	"testing"
	"time"

	"github.com/xinliangnote/go-gin-api/internal/api/third_party_request"
	"github.com/xinliangnote/go-gin-api/pkg/httpclient"
)

var demoGetAuthorization = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjEsIlVzZXJOYW1lIjoieGlubGlhbmdub3RlIiwiZXhwIjoxNjEzODI3MTEzLCJpYXQiOjE2MTM3NDA3MTMsIm5iZiI6MTYxMzc0MDcxM30.SnooP1ikO33ryGPdohsmOKqISa-bWzMkMvUNb5f2zc0"

func TestDemoGet(t *testing.T) {
	res, err := DemoGet("Tom",
		httpclient.WithTTL(time.Second*5),
		//httpclient.WithTrace(ctx.Trace()),
		//httpclient.WithLogger(ctx.Logger()),
		httpclient.WithHeader("Authorization", demoGetAuthorization),
		httpclient.WithOnFailedRetry(3, time.Second*1, DemoGetRetryVerify),
		httpclient.WithOnFailedAlarm("接口告警", new(third_party_request.AlarmEmail), DemoGetAlarmVerify),
		httpclient.WithMock(DemoGetMock),
	)

	if err != nil {
		t.Log("get [demo/get] err", err)
	}

	t.Log(res)
}
