package token

import (
	"net/url"
	"testing"
	"time"
)

const secret = "i1ydX9RtHyuJTrw7frcu"

func TestJwtSign(t *testing.T) {
	tokenString, err := New(secret).JwtSign(123456789, "xinliangnote", 24*time.Hour)
	if err != nil {
		t.Error("sign error", err)
		return
	}
	t.Log(tokenString)
}

func TestJwtParse(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjEyMzQ1Njc4OSwidXNlcm5hbWUiOiJ4aW5saWFuZyIsImV4cCI6MTYwOTQ2NzcwNCwiaWF0IjoxNjA5MzgxMzA0LCJpc3MiOiJnby1naW4tYXBpIiwibmJmIjoxNjA5MzgxMzA0fQ.hccv8F713WpKcwiSldBrFLZz_2SZzOTPedPi-8ps7M4"
	user, err := New(secret).JwtParse(tokenString)
	if err != nil {
		t.Error("parse error", err)
		return
	}
	t.Log(user)
}

func TestUrlSign(t *testing.T) {

	urlPath := "/echo"
	method := "post"
	params := url.Values{}
	params.Add("a", "a1")
	params.Add("d", "d1")
	params.Add("c", "c1")

	tokenString, err := New(secret).UrlSign(urlPath, method, params)
	if err != nil {
		t.Error("sign error", err)
		return
	}
	t.Log(tokenString)
}

func BenchmarkJwtSignAndParse(b *testing.B) {
	b.ResetTimer()
	token := New(secret)
	for i := 0; i < b.N; i++ {
		tokenString, _ := token.JwtSign(123456789, "xinliangnote", 24*time.Hour)
		token.JwtParse(tokenString)
	}
}
