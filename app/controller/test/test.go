package test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-api/app/util/aes"
	"go-gin-api/app/util/md5"
	"go-gin-api/app/util/response"
	"go-gin-api/app/util/rsa"
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
		sn, _ := aes.AesEncrypt(encryptStr, []byte(appSecret), appSecret)

		// 验证签名
		aes.AesDecrypt(sn, []byte(appSecret), appSecret)
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
		sn, _ := rsa.RsaPublicEncrypt(encryptStr, "rsa/public.pem")

		// 验证签名
		rsa.RsaPrivateDecrypt(sn, "rsa/private.pem")
	}
	utilGin := response.Gin{Ctx: c}
	utilGin.Response(1, fmt.Sprintf("%v次 - %v", count, time.Since(startTime)), nil)
}
