package cron

import (
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/validation"
	"github.com/xinliangnote/go-gin-api/internal/services/cron"

	"github.com/spf13/cast"
)

type detailRequest struct {
	Id string `uri:"id"` // HashID
}

type detailResponse struct {
	Name                string `json:"name"`                  // 任务名称
	Spec                string `json:"spec"`                  // crontab 表达式
	Command             string `json:"command"`               // 执行命令
	Protocol            int32  `json:"protocol"`              // 执行方式 1:shell 2:http
	HttpMethod          int32  `json:"http_method"`           // http 请求方式 1:get 2:post
	Timeout             int32  `json:"timeout"`               // 超时时间(单位:秒)
	RetryTimes          int32  `json:"retry_times"`           // 重试次数
	RetryInterval       int32  `json:"retry_interval"`        // 重试间隔(单位:秒)
	NotifyStatus        int32  `json:"notify_status"`         // 执行结束是否通知 1:不通知 2:失败通知 3:结束通知 4:结果关键字匹配通知
	NotifyType          int32  `json:"notify_type"`           // 通知类型 1:邮件 2:webhook
	NotifyReceiverEmail string `json:"notify_receiver_email"` // 通知者邮箱地址(多个用,分割)
	NotifyKeyword       string `json:"notify_keyword"`        // 通知匹配关键字(多个用,分割)
	Remark              string `json:"remark"`                // 备注
	IsUsed              int32  `json:"is_used"`               // 是否启用 1:是  -1:否
}

// Detail 获取单条任务详情
// @Summary 获取单条任务详情
// @Description 获取单条任务详情
// @Tags API.cron
// @Accept json
// @Produce json
// @Param id path string true "hashId"
// @Success 200 {object} detailResponse
// @Failure 400 {object} code.Failure
// @Router /api/cron/{id} [get]
// @Security LoginToken
func (h *handler) Detail() core.HandlerFunc {
	return func(ctx core.Context) {
		req := new(detailRequest)
		res := new(detailResponse)
		if err := ctx.ShouldBindURI(req); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				validation.Error(err)).WithError(err),
			)
			return
		}

		ids, err := h.hashids.HashidsDecode(req.Id)
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.HashIdsDecodeError,
				code.Text(code.HashIdsDecodeError)).WithError(err),
			)
			return
		}

		searchOneData := new(cron.SearchOneData)
		searchOneData.Id = cast.ToInt32(ids[0])

		info, err := h.cronService.Detail(ctx, searchOneData)
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.CronDetailError,
				code.Text(code.CronDetailError)).WithError(err),
			)
			return
		}

		res.Name = info.Name
		res.Spec = info.Spec
		res.Command = info.Command
		res.Protocol = info.Protocol
		res.HttpMethod = info.HttpMethod
		res.Timeout = info.Timeout
		res.RetryTimes = info.RetryTimes
		res.RetryInterval = info.RetryInterval
		res.NotifyStatus = info.NotifyStatus
		res.NotifyType = info.NotifyType
		res.NotifyReceiverEmail = info.NotifyReceiverEmail
		res.NotifyKeyword = info.NotifyKeyword
		res.Remark = info.Remark
		res.IsUsed = info.IsUsed

		ctx.Payload(res)
	}
}
