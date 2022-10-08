package cron_task

import "time"

// CronTask 后台任务表
//
//go:generate gormgen -structs CronTask -input .
type CronTask struct {
	Id                  int32     // 主键
	Name                string    // 任务名称
	Spec                string    // crontab 表达式
	Command             string    // 执行命令
	Protocol            int32     // 执行方式 1:shell 2:http
	HttpMethod          int32     // http 请求方式 1:get 2:post
	Timeout             int32     // 超时时间(单位:秒)
	RetryTimes          int32     // 重试次数
	RetryInterval       int32     // 重试间隔(单位:秒)
	NotifyStatus        int32     // 执行结束是否通知 1:不通知 2:失败通知 3:结束通知 4:结果关键字匹配通知
	NotifyType          int32     // 通知类型 1:邮件 2:webhook
	NotifyReceiverEmail string    // 通知者邮箱地址(多个用,分割)
	NotifyKeyword       string    // 通知匹配关键字(多个用,分割)
	Remark              string    // 备注
	IsUsed              int32     // 是否启用 1:是  -1:否
	CreatedAt           time.Time `gorm:"time"` // 创建时间
	CreatedUser         string    // 创建人
	UpdatedAt           time.Time `gorm:"time"` // 更新时间
	UpdatedUser         string    // 更新人
}
