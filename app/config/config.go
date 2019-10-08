package config

const (
	AppMode = "release" //debug or release
	AppPort = ":9999"
	AppName = "go-gin-api"

	// MD5 密钥
	AppMD5SignSecret = "4OhYXtDYNYxQsGetqASVOTP37jGt5gGY"

	// MD5 超时时间
	AppMD5SignExpiry = "120"

	// 超时时间
	AppReadTimeout  = 120
	AppWriteTimeout = 120

	// 日志文件
	AppAccessLogName = "log/" + AppName + "-access.log"
	AppErrorLogName  = "log/" + AppName + "-error.log"
	AppGrpcLogName   = "log/" + AppName + "-grpc.log"

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
	JaegerHostPort = "127.0.0.1:6831"

	// Jaeger 配置开关 1=开通 -1=关闭
	JaegerOpen = 1
)
