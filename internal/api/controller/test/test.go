package test

import (
	"fmt"
	"time"

	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/errno"

	"github.com/xinliangnote/go-util/aes"
	"github.com/xinliangnote/go-util/md5"
	"github.com/xinliangnote/go-util/rsa"
)

func Md5Test(c core.Context) {
	startTime := time.Now()
	appSecret := "IgkibX71IEf382PT"
	encryptStr := "param_1=xxx&param_2=xxx&ak=xxx&ts=1111111111"
	count := 1000000
	for i := 0; i < count; i++ {
		// 生成签名
		md5.MD5(appSecret + encryptStr + appSecret)

		// 验证签名
		md5.MD5(appSecret + encryptStr + appSecret)
	}
	c.SetPayload(errno.OK.WithData(fmt.Sprintf("%v次 - %v", count, time.Since(startTime))))
}

func AesTest(c core.Context) {
	startTime := time.Now()
	appSecret := "IgkibX71IEf382PT"
	encryptStr := "param_1=xxx&param_2=xxx&ak=xxx&ts=1111111111"
	count := 1000000
	for i := 0; i < count; i++ {
		// 生成签名
		sn, _ := aes.Encrypt(encryptStr, []byte(appSecret), appSecret)

		// 验证签名
		aes.Decrypt(sn, []byte(appSecret), appSecret)
	}
	c.SetPayload(errno.OK.WithData(fmt.Sprintf("%v次 - %v", count, time.Since(startTime))))
}

func RsaTest(c core.Context) {
	startTime := time.Now()
	encryptStr := "param_1=xxx&param_2=xxx&ak=xxx&ts=1111111111"
	count := 500
	for i := 0; i < count; i++ {
		// 生成签名
		sn, _ := rsa.PublicEncrypt(encryptStr, "rsa/public.pem")

		// 验证签名
		rsa.PrivateDecrypt(sn, "rsa/private.pem")
	}
	c.SetPayload(errno.OK.WithData(fmt.Sprintf("%v次 - %v", count, time.Since(startTime))))
}
