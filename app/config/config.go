package config

const (
	AppMode = "release" //debug or release
	AppPort = ":9999"
	AppName = "go-gin-api"

	// 超时时间
	AppReadTimeout  = 120
	AppWriteTimeout = 120

	// 日志文件
	AppAccessLogName = "log/" + AppName + "-access.log"
	AppErrorLogName  = "log/" + AppName + "-error.log"

	// 系统告警邮箱信息
	SystemEmailUser = "xinliangnote@163.com"
	SystemEmailPass = "" //密码或授权码
	SystemEmailHost = "smtp.163.com"
	SystemEmailPort = 465

	// 告警接收人
	ErrorNotifyUser = "xinliangnote@163.com"

	// 告警开关 1=开通 -1=关闭
	ErrorNotifyOpen = 1
)
