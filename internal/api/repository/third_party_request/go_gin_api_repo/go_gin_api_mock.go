package go_gin_api_repo

import "encoding/json"

func MockDemoGet() (body []byte) {
	res := new(demoGetResponse)
	res.Code = 1
	res.Msg = "ok"
	res.Data.Name = "AA"
	res.Data.Job = "AA_JOB"

	body, _ = json.Marshal(res)
	return body
}

func MockDemoPost() (body []byte) {
	res := new(demoPostResponse)
	res.Code = 1
	res.Msg = "ok"
	res.Data.Name = "BB"
	res.Data.Job = "BB_JOB"

	body, _ = json.Marshal(res)
	return body
}
