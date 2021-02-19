package go_gin_api_repo

import (
	"encoding/json"
	"net/url"

	"github.com/xinliangnote/go-gin-api/pkg/httpclient"

	"github.com/pkg/errors"
)

type demoGetResponse struct {
	Name string `json:"name"`
	Job  string `json:"job"`
}

type demoPostResponse struct {
	Name string `json:"name"`
	Job  string `json:"job"`
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

	return res, nil
}

func DemoGetRetryVerify(body []byte) (shouldRetry bool) {
	if len(body) == 0 {
		return true
	}

	return false
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

	return res, nil
}

func DemoPostRetryVerify(body []byte) (shouldRetry bool) {
	if len(body) == 0 {
		return true
	}

	return false
}
