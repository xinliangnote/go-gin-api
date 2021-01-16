package go_gin_api_repo

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/xinliangnote/go-gin-api/internal/api/repository/third_party_request"
	"github.com/xinliangnote/go-gin-api/pkg/httpclient"
)

var authorization = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjEwLCJVc2VyTmFtZSI6IjEyMyIsImV4cCI6MTYxMDcxODA3NCwiaWF0IjoxNjEwNjMxNjc0LCJpc3MiOiJnby1naW4tYXBpIiwibmJmIjoxNjEwNjMxNjc0fQ.S3T4MaIaz3XjkbJ-xkMDkwzuZ_jfZ8ZRf4cPMz0oXBE"

func TestDemoGet(t *testing.T) {
	res, err := DemoGet("Tom",
		httpclient.WithTTL(time.Second*5),
		//httpclient.WithTrace(ctx.Trace()),
		//httpclient.WithLogger(ctx.Logger()),
		httpclient.WithHeader("Authorization", authorization),
		httpclient.WithOnFailedRetry(3, time.Second*1, retryVerify),
		httpclient.WithOnFailedAlarm("接口告警", new(third_party_request.AlarmEmail), alarmVerify),
		//httpclient.WithMock(MockDemoGet),
	)

	if err != nil {
		t.Log("get [demo/get] err", err)
	}

	t.Log(res)
}

func TestDemoPost(t *testing.T) {
	res, err := DemoPost("Jack",
		httpclient.WithTTL(time.Second*5),
		//httpclient.WithTrace(ctx.Trace()),
		//httpclient.WithLogger(ctx.Logger()),
		httpclient.WithHeader("Authorization", authorization),
		httpclient.WithMock(MockDemoPost),
	)

	if err != nil {
		t.Log("post [demo/post] err", err)
	}

	t.Log(res)
}

// 设置重试规则
func retryVerify(body []byte) (shouldRetry bool) {
	if len(body) == 0 {
		return true
	}

	type Response struct {
		Code int `json:"code"`
	}
	resp := new(Response)
	if err := json.Unmarshal(body, resp); err != nil {
		return true
	}

	return resp.Code != 1
}

// 设置告警规则
func alarmVerify(body []byte) (shouldAlarm bool) {
	if len(body) == 0 {
		return true
	}

	type Response struct {
		Code int `json:"code"`
	}
	resp := new(Response)
	if err := json.Unmarshal(body, resp); err != nil {
		return true
	}

	return resp.Code != 1
}
