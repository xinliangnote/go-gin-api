package go_gin_api

import (
	"encoding/json"

	"github.com/xinliangnote/go-gin-api/pkg/httpclient"

	"github.com/pkg/errors"
)

// 接口地址
var demoGetApi = "http://127.0.0.1:9999/demo/get/"

// 接口返回结构
type demoGetResponse struct {
	Name string `json:"name"`
	Job  string `json:"job"`
}

// DemoGet 发起请求
func DemoGet(name string, opts ...httpclient.Option) (res *demoGetResponse, err error) {
	api := demoGetApi + name
	body, err := httpclient.Get(api, nil, opts...)
	if err != nil {
		return nil, err
	}

	res = new(demoGetResponse)
	err = json.Unmarshal(body, res)
	if err != nil {
		return nil, errors.Wrap(err, "DemoGet json unmarshal error")
	}

	return res, nil
}

// DemoGetRetryVerify 设置重试规则
func DemoGetRetryVerify(body []byte) (shouldRetry bool) {
	if len(body) == 0 {
		return true
	}

	return false
}

// DemoGetAlarmVerify 设置告警规则
func DemoGetAlarmVerify(body []byte) (shouldAlarm bool) {
	if len(body) == 0 {
		return true
	}

	return false
}

// DemoGetMock 设置 Mock 数据
func DemoGetMock() (body []byte) {
	res := new(demoGetResponse)
	res.Name = "AA"
	res.Job = "AA_JOB"

	body, _ = json.Marshal(res)
	return body
}
