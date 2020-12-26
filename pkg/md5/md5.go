package md5

import (
	cryptoMD5 "crypto/md5"
	"encoding/hex"
)

var _ MD5 = (*md5)(nil)

type MD5 interface {
	i()
	// Encrypt 加密
	Encrypt(encryptStr string) string
}

type md5 struct{}

func New() MD5 {
	return &md5{}
}

func (m *md5) i() {}

func (m *md5) Encrypt(encryptStr string) string {
	s := cryptoMD5.New()
	s.Write([]byte(encryptStr))
	return hex.EncodeToString(s.Sum(nil))
}
