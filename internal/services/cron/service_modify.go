package cron

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql/cron_task"

	"github.com/spf13/cast"
)

type ModifyCronTaskData struct {
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

func (s *service) Modify(ctx core.Context, id int32, modifyData *ModifyCronTaskData) (err error) {
	data := map[string]interface{}{
		"name":                  modifyData.Name,
		"spec":                  modifyData.Spec,
		"command":               modifyData.Command,
		"protocol":              modifyData.Protocol,
		"http_method":           modifyData.HttpMethod,
		"timeout":               modifyData.Timeout,
		"retry_times":           modifyData.RetryTimes,
		"retry_interval":        modifyData.RetryInterval,
		"notify_status":         modifyData.NotifyStatus,
		"notify_type":           modifyData.NotifyType,
		"notify_receiver_email": modifyData.NotifyReceiverEmail,
		"notify_keyword":        modifyData.NotifyKeyword,
		"remark":                modifyData.Remark,
		"is_used":               modifyData.IsUsed,
		"updated_user":          ctx.SessionUserInfo().UserName,
	}

	qb := cron_task.NewQueryBuilder()
	qb.WhereId(mysql.EqualPredicate, id)
	err = qb.Updates(s.db.GetDbW().WithContext(ctx.RequestContext()), data)
	if err != nil {
		return err
	}

	// region 操作定时任务 避免主从同步延迟，在这需要查询主库
	if modifyData.IsUsed == cron_task.IsUsedNo {
		s.cronServer.RemoveTask(cast.ToInt(id))
	} else {
		qb = cron_task.NewQueryBuilder()
		qb.WhereId(mysql.EqualPredicate, id)
		info, err := qb.QueryOne(s.db.GetDbW().WithContext(ctx.RequestContext()))
		if err != nil {
			return err
		}

		s.cronServer.RemoveTask(cast.ToInt(id))
		s.cronServer.AddTask(info)
	}
	// endregion

	return
}
