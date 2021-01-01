package md5

import "testing"

func TestEncrypt(t *testing.T) {
	t.Log(New().Encrypt("123456"))
}
