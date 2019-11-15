package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"os"
)

// 公钥加密
func PublicEncrypt(encryptStr string, path string) (string, error) {
	// 打开文件
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// 读取文件内容
	info, _ := file.Stat()
	buf := make([]byte,info.Size())
	file.Read(buf)

	// pem 解码
	block, _ := pem.Decode(buf)

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

// 私钥解密
func PrivateDecrypt(decryptStr string, path string) (string, error) {
	// 打开文件
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// 获取文件内容
	info, _ := file.Stat()
	buf := make([]byte,info.Size())
	file.Read(buf)

	// pem 解码
	block, _ := pem.Decode(buf)

	// X509 解码
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	decryptBytes, err := base64.URLEncoding.DecodeString(decryptStr)

	//对密文进行解密
	decrypted, _ := rsa.DecryptPKCS1v15(rand.Reader,privateKey,decryptBytes)

	//返回明文
	return string(decrypted), nil
}

