package signature

import (
	"net/url"
	"testing"
	"time"
)

const (
	key    = "blog"
	secret = "i1ydX9RtHyuJTrw7frcu"
	ttl    = time.Minute * 10
)

func TestSignature_Generate(t *testing.T) {
	path := "/echo"
	method := "POST"

	params := url.Values{}
	params.Add("a", "a1")
	params.Add("d", "d1")
	params.Add("c", "c1")

	authorization, date, err := New(key, secret, ttl).Generate(path, method, params)
	t.Log("authorization:", authorization)
	t.Log("date:", date)
	t.Log("err:", err)
}

func TestSignature_Verify(t *testing.T) {

	authorization := "blog d3+XC7l0lproeO4Z/WHd2VWZsBct0dy9HUndNOxmPu0="
	date := "Fri, 02 Apr 2021 20:31:28 GMT"

	path := "/echo"
	method := "post"
	params := url.Values{}
	params.Add("a", "a1")
	params.Add("d", "d1")
	params.Add("c", "c1")

	ok, err := New(key, secret, ttl).Verify(authorization, date, path, method, params)
	t.Log(ok)
	t.Log(err)
}
