package aes

import "testing"

const (
	Key = "IgkibX71IEf382PT"
	Iv  = "IgkibX71IEf382PT"
)

func TestEncrypt(t *testing.T) {
	t.Log(New(Key, Iv).Encrypt("123456"))
}

func TestDecrypt(t *testing.T) {
	t.Log(New(Key, Iv).Decrypt("GO-ri84zevE-z1biJwfQPw=="))
}
