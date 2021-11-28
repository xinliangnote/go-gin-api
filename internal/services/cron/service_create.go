package cron

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql/cron_task"
)

type CreateCronTaskData struct {
	Name                string // 任务名称
	Spec                string // crontab 表达式
	Command             string // 执行命令
	Protocol            int32  // 执行方式 1:shell 2:http
	HttpMethod          int32  // http 请求方式 1:get 2:post
	Timeout             int32  // 超时时间(单位:秒)
	RetryTimes          int32  // 重试次数
	RetryInterval       int32  // 重试间隔(单位:秒)
	NotifyStatus        int32  // 执行结束是否通知 1:不通知 2:失败通知 3:结束通知 4:结果关键字匹配通知
	NotifyType          int32  // 通知类型 1:邮件 2:webhook
	NotifyReceiverEmail string // 通知者邮箱地址(多个用,分割)
	NotifyKeyword       string // 通知匹配关键字(多个用,分割)
	Remark              string // 备注
	IsUsed              int32  // 是否启用 1:是  -1:否
}

func (s *service) Create(ctx core.Context, createData *CreateCronTaskData) (id int32, err error) {
	model := cron_task.NewModel()
	model.Name = createData.Name
	model.Spec = createData.Spec
	model.Command = createData.Command
	model.Protocol = createData.Protocol
	model.HttpMethod = createData.HttpMethod
	model.Timeout = createData.Timeout
	model.RetryTimes = createData.RetryTimes
	model.RetryInterval = createData.RetryInterval
	model.NotifyStatus = createData.NotifyStatus
	model.NotifyType = createData.NotifyType
	model.NotifyReceiverEmail = createData.NotifyReceiverEmail
	model.NotifyKeyword = createData.NotifyKeyword
	model.Remark = createData.Remark
	model.IsUsed = createData.IsUsed
	model.CreatedUser = ctx.SessionUserInfo().UserName

	id, err = model.Create(s.db.GetDbW().WithContext(ctx.RequestContext()))
	if err != nil {
		return 0, err
	}

	s.cronServer.AddTask(model)

	return
}
