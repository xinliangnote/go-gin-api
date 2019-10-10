package test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-api/app/util"
	"time"
)

func Md5Test(c *gin.Context) {
	startTime  := time.Now()
	appSecret  := "IgkibX71IEf382PT"
	encryptStr := "param_1=xxx&param_2=xxx&ak=xxx&ts=1111111111"
	count      := 1000000
	for i := 0; i < count; i++ {
		// 生成签名
		util.MD5(appSecret + encryptStr + appSecret)

		// 验证签名
		util.MD5(appSecret + encryptStr + appSecret)
	}
	utilGin := util.Gin{Ctx: c}
	utilGin.Response(1, fmt.Sprintf("%v次 - %v", count, time.Since(startTime)), nil)
}

func AesTest(c *gin.Context) {
	startTime  := time.Now()
	appSecret  := "IgkibX71IEf382PT"
	encryptStr := "param_1=xxx&param_2=xxx&ak=xxx&ts=1111111111"
	count      := 1000000
	for i := 0; i < count; i++ {
		// 生成签名
		sn, _ := util.AesEncrypt(encryptStr, []byte(appSecret), appSecret)

		// 验证签名
		util.AesDecrypt(sn, []byte(appSecret), appSecret)
	}
	utilGin := util.Gin{Ctx: c}
	utilGin.Response(1, fmt.Sprintf("%v次 - %v", count, time.Since(startTime)), nil)
}

func RsaTest(c *gin.Context) {
	startTime  := time.Now()
	encryptStr := "param_1=xxx&param_2=xxx&ak=xxx&ts=1111111111"
	count      := 500
	for i := 0; i < count; i++ {
		// 生成签名
		sn, _ := util.RsaPublicEncrypt(encryptStr, "rsa/public.pem")

		// 验证签名
		util.RsaPrivateDecrypt(sn, "rsa/private.pem")
	}
	utilGin := util.Gin{Ctx: c}
	utilGin.Response(1, fmt.Sprintf("%v次 - %v", count, time.Since(startTime)), nil)
}
