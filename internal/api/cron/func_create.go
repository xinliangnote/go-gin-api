package cron

import (
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/validation"
	"github.com/xinliangnote/go-gin-api/internal/services/cron"
)

type createRequest struct {
	Name                string `form:"name" binding:"required"`           // 任务名称
	Spec                string `form:"spec" binding:"required"`           // crontab 表达式
	Command             string `form:"command" binding:"required"`        // 执行命令
	Protocol            int32  `form:"protocol" binding:"required"`       // 执行方式 1:shell 2:http
	HttpMethod          int32  `form:"http_method"`                       // http 请求方式 1:get 2:post
	Timeout             int32  `form:"timeout" binding:"required"`        // 超时时间(单位:秒)
	RetryTimes          int32  `form:"retry_times" binding:"required"`    // 重试次数
	RetryInterval       int32  `form:"retry_interval" binding:"required"` // 重试间隔(单位:秒)
	NotifyStatus        int32  `form:"notify_status" binding:"required"`  // 执行结束是否通知 1:不通知 2:失败通知 3:结束通知 4:结果关键字匹配通知
	NotifyType          int32  `form:"notify_type"`                       // 通知类型 1:邮件 2:webhook
	NotifyReceiverEmail string `form:"notify_receiver_email"`             // 通知者邮箱地址(多个用,分割)
	NotifyKeyword       string `form:"notify_keyword"`                    // 通知匹配关键字(多个用,分割)
	Remark              string `form:"remark"`                            // 备注
	IsUsed              int32  `form:"is_used" binding:"required"`        // 是否启用 1:是  -1:否
}

type createResponse struct {
	Id int32 `json:"id"` // 主键ID
}

// Create 创建任务
// @Summary 创建任务
// @Description 创建任务
// @Tags API.cron
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param name formData string true "任务名称"
// @Param spec formData string true "crontab 表达式"
// @Param command formData string true "执行命令"
// @Param protocol formData int true "执行方式 1:shell 2:http"
// @Param http_method formData int false "http 请求方式 1:get 2:post"
// @Param timeout formData int true "超时时间(单位:秒)"
// @Param retry_times formData int true "重试次数"
// @Param retry_interval formData int true "重试间隔(单位:秒)"
// @Param notify_status formData int true "执行结束是否通知 1:不通知 2:失败通知 3:结束通知 4:结果关键字匹配通知"
// @Param notify_type formData int false "通知类型 1:邮件 2:webhook"
// @Param notify_receiver_email formData string false "通知者邮箱地址(多个用,分割)"
// @Param notify_keyword formData string false "通知匹配关键字(多个用,分割)"
// @Param remark formData string false "备注"
// @Param is_used formData int true "是否启用 1:是  -1:否"
// @Success 200 {object} createResponse
// @Failure 400 {object} code.Failure
// @Router /api/cron [post]
// @Security LoginToken
func (h *handler) Create() core.HandlerFunc {
	return func(ctx core.Context) {
		req := new(createRequest)
		res := new(createResponse)
		if err := ctx.ShouldBindForm(req); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				validation.Error(err)).WithError(err),
			)
			return
		}

		createData := new(cron.CreateCronTaskData)
		createData.Name = req.Name
		createData.Spec = req.Spec
		createData.Command = req.Command
		createData.Protocol = req.Protocol
		createData.HttpMethod = req.HttpMethod
		createData.Timeout = req.Timeout
		createData.RetryTimes = req.RetryTimes
		createData.RetryInterval = req.RetryInterval
		createData.NotifyStatus = req.NotifyStatus
		createData.NotifyType = req.NotifyType
		createData.NotifyReceiverEmail = req.NotifyReceiverEmail
		createData.NotifyKeyword = req.NotifyKeyword
		createData.Remark = req.Remark
		createData.IsUsed = req.IsUsed

		id, err := h.cronService.Create(ctx, createData)
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.CronCreateError,
				code.Text(code.CronCreateError)).WithError(err),
			)
			return
		}

		res.Id = id
		ctx.Payload(res)
	}
}
