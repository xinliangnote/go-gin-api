package token

import (
	"net/url"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var _ Token = (*token)(nil)

type Token interface {
	// i 为了避免被其他包实现
	i()

	// JwtSign 签名
	JwtSign(userId int64, userName string, expireDuration time.Duration) (tokenString string, err error)

	// JwtParse 解密
	JwtParse(tokenString string) (*claims, error)

	// UrlSign URL 签名方式，不支持解密
	UrlSign(path string, method string, params url.Values) (tokenString string, err error)
}

type token struct {
	secret string
}

type claims struct {
	UserID   int64
	UserName string
	jwt.StandardClaims
}

func New(secret string) Token {
	return &token{
		secret: secret,
	}
}

func (t *token) i() {}
