package configs

const (
	// 项目版本
	ProjectVersion = "v1.2.6"

	// 项目名称
	ProjectName = "go-gin-api"

	// 项目端口
	ProjectPort = ":9999"

	// 项目日志存放文件
	ProjectLogFile = "./logs/" + ProjectName + "-access.log"

	// 项目安装完成标识
	ProjectInstallMark = "INSTALL.lock"

	// 登录验证 Token，Header 中传递的参数
	LoginToken = "Token"

	// 签名验证 Token，Header 中传递的参数
	SignToken = "Authorization"

	// 签名验证 Date，Header 中传递的参数
	SignTokenDate = "Authorization-Date"

	// Redis Key 前缀 - 防止重复提交
	RedisKeyPrefixRequestID = ProjectName + ":request-id:"

	// Redis Key 前缀 - 登录用户信息
	RedisKeyPrefixLoginUser = ProjectName + ":login-user:"

	// Redis Key 前缀 - 签名验证信息
	RedisKeyPrefixSignature = ProjectName + ":signature:"
)
