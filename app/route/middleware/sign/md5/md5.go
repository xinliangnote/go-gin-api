package sign_md5

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-api/app/config"
	"go-gin-api/app/util"
	"net/url"
	"sort"
	"strconv"
	"time"
)

// MD5 组合加密
func SetUp() gin.HandlerFunc {

	return func(c *gin.Context) {
		utilGin := util.Gin{Ctx: c}

		sign, err := verifyMD5Sign(c)

		if sign != nil {
			utilGin.Response(-1, "Debug Sign", sign)
			c.Abort()
			return
		}

		if err != nil {
			utilGin.Response(-1, err.Error(), sign)
			c.Abort()
			return
		}

		c.Next()
	}
}

// 创建签名
func createMD5Sign(params url.Values) string {
	var key []string
	var str = ""
	for k := range params {
		if k != "sn" && k != "ts" && k != "debug" {
			key = append(key, k)
		}
	}
	sort.Strings(key)
	for i := 0; i < len(key); i++ {
		if i == 0 {
			str = fmt.Sprintf("%v=%v", key[i], params.Get(key[i]))
		} else {
			str = str + fmt.Sprintf("&%v=%v", key[i], params.Get(key[i]))
		}
	}

	// 自定义签名算法
	sign := util.MD5(config.AppSignSecret + str + config.AppSignSecret)
	return sign
}

// 验证签名
func verifyMD5Sign(c *gin.Context) (map[string]string, error) {
	var method = c.Request.Method
	var ts int64
	var sn string
	var req url.Values
	var debug string

	if method == "GET" {
		req = c.Request.URL.Query()
		sn = c.Query("sn")
		debug = c.Query("debug")
		ts, _ = strconv.ParseInt(c.Query("ts"), 10, 64)
	} else if method == "POST" {
		_ = c.Request.ParseForm()
		req = c.Request.PostForm
		sn = c.PostForm("sn")
		debug = c.PostForm("debug")
		ts, _ = strconv.ParseInt(c.PostForm("ts"), 10, 64)
	} else {
		return nil, errors.New("非法请求")
	}

	if debug == "1" {
		res := map[string]string{
			"ts": strconv.FormatInt(util.GetCurrentUnix(), 10),
			"sn": createMD5Sign(req),
		}
		return res, nil
	}

	exp, _ := strconv.ParseInt(config.AppSignExpiry, 10, 64)

	// 验证过期时间
	timestamp := time.Now().Unix()
	if ts > timestamp || timestamp - ts >= exp {
		return nil, errors.New("ts Error")
	}

	// 验证签名
	if sn == "" || sn != createMD5Sign(req) {
		return nil, errors.New("sn Error")
	}

	return nil, nil
}
