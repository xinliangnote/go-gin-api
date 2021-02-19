package go_gin_api_repo

import "encoding/json"

func MockDemoGet() (body []byte) {
	res := new(demoGetResponse)
	res.Name = "AA"
	res.Job = "AA_JOB"

	body, _ = json.Marshal(res)
	return body
}

func MockDemoPost() (body []byte) {
	res := new(demoPostResponse)
	res.Name = "BB"
	res.Job = "BB_JOB"

	body, _ = json.Marshal(res)
	return body
}
