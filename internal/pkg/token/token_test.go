package token

import (
	"testing"
)

// 执行 Test 时，先将 token.secret 设置值

func TestSign(t *testing.T) {
	tokenString, err := Sign(123456789, "xinliangnote")
	if err != nil {
		t.Error("sign error", err)
		return
	}
	t.Log(tokenString)
}

func TestParse(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjEyMzQ1Njc4OSwidXNlcm5hbWUiOiJ4aW5saWFuZyIsImV4cCI6MTYwOTQ2NzcwNCwiaWF0IjoxNjA5MzgxMzA0LCJpc3MiOiJnby1naW4tYXBpIiwibmJmIjoxNjA5MzgxMzA0fQ.hccv8F713WpKcwiSldBrFLZz_2SZzOTPedPi-8ps7M4"
	user, err := Parse(tokenString)
	if err != nil {
		t.Error("parse error", err)
		return
	}
	t.Log(user)
}
