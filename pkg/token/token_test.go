package token

import (
	"testing"
	"time"
)

const secret = "i1ydX9RtHyuJTrw7frcu"

func TestSign(t *testing.T) {
	tokenString, err := New(secret).Sign(123456789, "xinliangnote", 24*time.Hour)
	if err != nil {
		t.Error("sign error", err)
		return
	}
	t.Log(tokenString)
}

func TestParse(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjEyMzQ1Njc4OSwidXNlcm5hbWUiOiJ4aW5saWFuZyIsImV4cCI6MTYwOTQ2NzcwNCwiaWF0IjoxNjA5MzgxMzA0LCJpc3MiOiJnby1naW4tYXBpIiwibmJmIjoxNjA5MzgxMzA0fQ.hccv8F713WpKcwiSldBrFLZz_2SZzOTPedPi-8ps7M4"
	user, err := New(secret).Parse(tokenString)
	if err != nil {
		t.Error("parse error", err)
		return
	}
	t.Log(user)
}

func BenchmarkSignAndParse(b *testing.B) {
	b.ResetTimer()
	token := New(secret)
	for i := 0; i < b.N; i++ {
		tokenString, _ := token.Sign(123456789, "xinliangnote", 24*time.Hour)
		token.Parse(tokenString)
	}
}
