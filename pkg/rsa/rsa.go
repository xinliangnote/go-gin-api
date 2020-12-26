package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
)

var _ Public = (*rsaPub)(nil)
var _ Private = (*rsaPri)(nil)

type Public interface {
	i()
	// Encrypt 加密
	Encrypt(encryptStr string) (string, error)
}

type Private interface {
	i()
	// Decrypt 解密
	Decrypt(decryptStr string) (string, error)
}

type rsaPub struct {
	PublicKey string
}

type rsaPri struct {
	PrivateKey string
}

func NewPublic(publicKey string) Public {
	return &rsaPub{
		PublicKey: publicKey,
	}
}

func NewPrivate(privateKey string) Private {
	return &rsaPri{
		PrivateKey: privateKey,
	}
}

func (pub *rsaPub) i() {}

func (pub *rsaPub) Encrypt(encryptStr string) (string, error) {
	// pem 解码
	block, _ := pem.Decode([]byte(pub.PublicKey))

	// x509 解码
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", err
	}

	// 类型断言
	publicKey := publicKeyInterface.(*rsa.PublicKey)

	//对明文进行加密
	encryptedStr, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(encryptStr))
	if err != nil {
		return "", err
	}

	//返回密文
	return base64.URLEncoding.EncodeToString(encryptedStr), nil
}

func (pri *rsaPri) i() {}

func (pri *rsaPri) Decrypt(decryptStr string) (string, error) {
	// pem 解码
	block, _ := pem.Decode([]byte(pri.PrivateKey))

	// X509 解码
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	decryptBytes, err := base64.URLEncoding.DecodeString(decryptStr)

	//对密文进行解密
	decrypted, _ := rsa.DecryptPKCS1v15(rand.Reader, privateKey, decryptBytes)

	//返回明文
	return string(decrypted), nil
}
