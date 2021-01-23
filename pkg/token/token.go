package token

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var _ Token = (*token)(nil)

type Token interface {
	// i 为了避免被其他包实现
	i()
	Sign(userId int64, userName string, expireDuration time.Duration) (tokenString string, err error)
	Parse(tokenString string) (*claims, error)
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

func (t *token) Sign(userId int64, userName string, expireDuration time.Duration) (tokenString string, err error) {
	// The token content.
	// iss: （Issuer）签发者
	// iat: （Issued At）签发时间，用Unix时间戳表示
	// exp: （Expiration Time）过期时间，用Unix时间戳表示
	// aud: （Audience）接收该JWT的一方
	// sub: （Subject）该JWT的主题
	// nbf: （Not Before）不要早于这个时间
	// jti: （JWT ID）用于标识JWT的唯一ID
	claims := claims{
		userId,
		userName,
		jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(expireDuration).Unix(),
		},
	}
	tokenString, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(t.secret))
	return
}

func (t *token) Parse(tokenString string) (*claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(tokenString, &claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.secret), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
