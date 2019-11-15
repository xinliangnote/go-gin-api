package test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xinliangnote/go-util/aes"
	"github.com/xinliangnote/go-util/md5"
	"github.com/xinliangnote/go-util/rsa"
	"go-gin-api/app/util/response"
	"time"
)

func Md5Test(c *gin.Context) {
	startTime  := time.Now()
	appSecret  := "IgkibX71IEf382PT"
	encryptStr := "param_1=xxx&param_2=xxx&ak=xxx&ts=1111111111"
	count      := 1000000
	for i := 0; i < count; i++ {
		// 生成签名
		md5.MD5(appSecret + encryptStr + appSecret)

		// 验证签名
		md5.MD5(appSecret + encryptStr + appSecret)
	}
	utilGin := response.Gin{Ctx: c}
	utilGin.Response(1, fmt.Sprintf("%v次 - %v", count, time.Since(startTime)), nil)
}

func AesTest(c *gin.Context) {
	startTime  := time.Now()
	appSecret  := "IgkibX71IEf382PT"
	encryptStr := "param_1=xxx&param_2=xxx&ak=xxx&ts=1111111111"
	count      := 1000000
	for i := 0; i < count; i++ {
		// 生成签名
		sn, _ := aes.Encrypt(encryptStr, []byte(appSecret), appSecret)

		// 验证签名
		aes.Decrypt(sn, []byte(appSecret), appSecret)
	}
	utilGin := response.Gin{Ctx: c}
	utilGin.Response(1, fmt.Sprintf("%v次 - %v", count, time.Since(startTime)), nil)
}

func RsaTest(c *gin.Context) {
	startTime  := time.Now()
	encryptStr := "param_1=xxx&param_2=xxx&ak=xxx&ts=1111111111"
	count      := 500
	for i := 0; i < count; i++ {
		// 生成签名
		sn, _ := rsa.PublicEncrypt(encryptStr, "rsa/public.pem")

		// 验证签名
		rsa.PrivateDecrypt(sn, "rsa/private.pem")
	}
	utilGin := response.Gin{Ctx: c}
	utilGin.Response(1, fmt.Sprintf("%v次 - %v", count, time.Since(startTime)), nil)
}
