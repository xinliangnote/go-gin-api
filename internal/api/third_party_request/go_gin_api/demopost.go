package go_gin_api

import (
	"encoding/json"
	"net/url"

	"github.com/xinliangnote/go-gin-api/pkg/httpclient"

	"github.com/pkg/errors"
)

// 接口地址
var demoPostApi = "http://127.0.0.1:9999/demo/post/"

// 接口返回结构
type demoPostResponse struct {
	Name string `json:"name"`
	Job  string `json:"job"`
}

// DemoPost 发起请求
func DemoPost(name string, opts ...httpclient.Option) (res *demoPostResponse, err error) {
	api := demoPostApi
	params := url.Values{}
	params.Set("name", name)
	body, err := httpclient.PostForm(api, params, opts...)
	if err != nil {
		return nil, err
	}

	res = new(demoPostResponse)
	err = json.Unmarshal(body, res)
	if err != nil {
		return nil, errors.Wrap(err, "DemoPost json unmarshal error")
	}

	return res, nil
}

// DemoPostRetryVerify 设置重试规则
func DemoPostRetryVerify(body []byte) (shouldRetry bool) {
	if len(body) == 0 {
		return true
	}

	return false
}

// DemoPostAlarmVerify 设置告警规则
func DemoPostAlarmVerify(body []byte) (shouldAlarm bool) {
	if len(body) == 0 {
		return true
	}

	return false
}

// DemoPostMock 设置 Mock 数据
func DemoPostMock() (body []byte) {
	res := new(demoPostResponse)
	res.Name = "BB"
	res.Job = "BB_JOB"

	body, _ = json.Marshal(res)
	return body
}
