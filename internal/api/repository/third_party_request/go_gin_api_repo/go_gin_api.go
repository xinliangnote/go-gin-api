package go_gin_api_repo

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/xinliangnote/go-gin-api/pkg/httpclient"

	"github.com/pkg/errors"
)

type demoGetResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Name string `json:"name"`
		Job  string `json:"job"`
	} `json:"data"`
	ID string `json:"id"`
}

type demoPostResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Name string `json:"name"`
		Job  string `json:"job"`
	} `json:"data"`
	ID string `json:"id"`
}

func DemoGet(name string, opts ...httpclient.Option) (res *demoGetResponse, err error) {
	api := "http://127.0.0.1:9999/demo/get/" + name
	body, err := httpclient.Get(api, nil, opts...)
	if err != nil {
		return nil, err
	}

	res = new(demoGetResponse)
	err = json.Unmarshal(body, res)
	if err != nil {
		return nil, errors.Wrap(err, "DemoGet json unmarshal error")
	}

	if res.Code != 1 {
		return nil, errors.New(fmt.Sprintf("code err: %d-%s", res.Code, res.Msg))
	}

	return res, nil
}

func DemoGetRetryVerify(body []byte) (shouldRetry bool) {
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

	// 例如 无需重试的 code 码，code !=1 需要重试
	successCode := 1
	return resp.Code != successCode
}

func DemoPost(name string, opts ...httpclient.Option) (res *demoPostResponse, err error) {
	api := "http://127.0.0.1:9999/demo/post"
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

	if res.Code != 1 {
		return nil, errors.New(fmt.Sprintf("code err: %d-%s", res.Code, res.Msg))
	}

	return res, nil
}

func DemoPostRetryVerify(body []byte) (shouldRetry bool) {
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

	// 例如 无需重试的 code 码，code !=1 需要重试
	successCode := 1
	return resp.Code != successCode
}
