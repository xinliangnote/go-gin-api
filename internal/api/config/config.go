package config

var (
	ApiAuthConfig = map[string]map[string]string{

		// 调用方
		"DEMO": {
			"md5": "IgkibX71IEf382PT",
			"aes": "IgkibX71IEf382PT",
			"rsa": "rsa/public.pem",
		},
	}
)

const (
	AppName = "go-gin-api"

	// 签名超时时间
	AppSignExpiry = "120"

	// RSA Private File
	AppRsaPrivateFile = "rsa/private.pem"

	// 日志文件
	AppErrorLogName = "log/" + AppName + "-error.log"
	AppGrpcLogName  = "log/" + AppName + "-grpc.log"

	// 系统告警邮箱信息
	SystemEmailUser = "xinliangnote@163.com"
	SystemEmailPass = "" //密码或授权码
	SystemEmailHost = "smtp.163.com"
	SystemEmailPort = 465

	// 告警接收人
	ErrorNotifyUser = "xinliangnote@163.com"

	// 告警开关 1=开通 -1=关闭
	ErrorNotifyOpen = -1

	// Jaeger 配置信息
	JaegerHostPort = "172.20.2.212:6831"

	// Jaeger 配置开关 1=开通 -1=关闭
	JaegerOpen = 1
)
