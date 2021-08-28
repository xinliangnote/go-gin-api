package cron_handler

import (
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/cron_task_repo"
	"github.com/xinliangnote/go-gin-api/internal/api/service/cron_service"
	"github.com/xinliangnote/go-gin-api/internal/pkg/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/validation"
	"github.com/xinliangnote/go-gin-api/pkg/errno"
	"github.com/xinliangnote/go-gin-api/pkg/time_parse"

	"github.com/spf13/cast"
)

type listRequest struct {
	Page     int    `form:"page"`      // 第几页
	PageSize int    `form:"page_size"` // 每页显示条数
	Name     string `form:"name"`      // 任务名称
	Protocol int    `form:"protocol"`  // 执行方式 1:shell 2:http
	IsUsed   int    `form:"is_used"`   // 是否启用 1:是  -1:否
}

type listData struct {
	Id               int    `json:"id"`                 // ID
	HashID           string `json:"hashid"`             // hashid
	Name             string `json:"name"`               // 任务名称
	Protocol         int    `json:"protocol"`           // 执行方式 1:shell 2:http
	ProtocolText     string `json:"protocol_text"`      // 执行方式
	Spec             string `json:"spec"`               // crontab 表达式
	Command          string `json:"command"`            // 执行命令
	HttpMethod       int    `json:"http_method"`        // http 请求方式 1:get 2:post
	HttpMethodText   string `json:"http_method_text"`   // http 请求方式
	Timeout          int    `json:"timeout"`            // 超时时间(单位:秒)
	RetryTimes       int    `json:"retry_times"`        // 重试次数
	RetryInterval    int    `json:"retry_interval"`     // 重试间隔(单位:秒)
	NotifyStatus     int    `json:"notify_status"`      // 执行结束是否通知 1:不通知 2:失败通知 3:结束通知 4:结果关键字匹配通知
	NotifyStatusText string `json:"notify_status_text"` // 执行结束是否通知
	IsUsed           int    `json:"is_used"`            // 是否启用 1=启用 2=禁用
	IsUsedText       string `json:"is_used_text"`       // 是否启用
	CreatedAt        string `json:"created_at"`         // 创建时间
	CreatedUser      string `json:"created_user"`       // 创建人
	UpdatedAt        string `json:"updated_at"`         // 更新时间
	UpdatedUser      string `json:"updated_user"`       // 更新人
}

type listResponse struct {
	List       []listData `json:"list"`
	Pagination struct {
		Total        int `json:"total"`
		CurrentPage  int `json:"current_page"`
		PerPageCount int `json:"per_page_count"`
	} `json:"pagination"`
}

// List 任务列表
// @Summary 任务列表
// @Description 任务列表
// @Tags API.cron
// @Accept multipart/form-data
// @Produce json
// @Param page query int false "第几页"
// @Param page_size query string false "每页显示条数"
// @Param name query string false "任务名称"
// @Param protocol query int false "执行方式 1:shell 2:http"
// @Param is_used query int false "是否启用 1:是  -1:否"
// @Success 200 {object} listResponse
// @Failure 400 {object} code.Failure
// @Router /api/cron [get]
func (h *handler) List() core.HandlerFunc {
	return func(ctx core.Context) {
		req := new(listRequest)
		res := new(listResponse)
		if err := ctx.ShouldBindForm(req); err != nil {
			ctx.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ParamBindError,
				validation.Error(err)).WithErr(err),
			)
			return
		}

		page := req.Page
		if page == 0 {
			page = 1
		}

		pageSize := req.PageSize
		if pageSize == 0 {
			pageSize = 10
		}

		searchData := new(cron_service.SearchData)
		searchData.Page = req.Page
		searchData.PageSize = req.PageSize
		searchData.Name = req.Name
		searchData.Protocol = cast.ToInt32(req.Protocol)
		searchData.IsUsed = cast.ToInt32(req.IsUsed)

		resListData, err := h.cronService.PageList(ctx, searchData)
		if err != nil {
			ctx.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.CronListError,
				code.Text(code.CronListError)).WithErr(err),
			)
			return
		}

		resCountData, err := h.cronService.PageListCount(ctx, searchData)
		if err != nil {
			ctx.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.CronListError,
				code.Text(code.CronListError)).WithErr(err),
			)
			return
		}

		res.Pagination.Total = cast.ToInt(resCountData)
		res.Pagination.PerPageCount = pageSize
		res.Pagination.CurrentPage = page
		res.List = make([]listData, len(resListData))

		for k, v := range resListData {
			hashId, err := h.hashids.HashidsEncode([]int{cast.ToInt(v.Id)})
			if err != nil {
				ctx.AbortWithError(errno.NewError(
					http.StatusBadRequest,
					code.HashIdsEncodeError,
					code.Text(code.HashIdsEncodeError)).WithErr(err),
				)
				return
			}

			data := listData{
				Id:               cast.ToInt(v.Id),
				HashID:           hashId,
				Name:             v.Name,
				Protocol:         cast.ToInt(v.Protocol),
				ProtocolText:     cron_task_repo.ProtocolText[v.Protocol],
				Spec:             v.Spec,
				Command:          v.Command,
				HttpMethod:       cast.ToInt(v.HttpMethod),
				HttpMethodText:   cron_task_repo.HttpMethodText[v.HttpMethod],
				Timeout:          cast.ToInt(v.Timeout),
				RetryTimes:       cast.ToInt(v.RetryTimes),
				RetryInterval:    cast.ToInt(v.RetryInterval),
				NotifyStatus:     cast.ToInt(v.NotifyStatus),
				NotifyStatusText: cron_task_repo.NotifyStatusText[v.NotifyStatus],
				IsUsed:           cast.ToInt(v.IsUsed),
				IsUsedText:       cron_task_repo.IsUsedText[v.IsUsed],
				CreatedAt:        v.CreatedAt.Format(time_parse.CSTLayout),
				CreatedUser:      v.CreatedUser,
				UpdatedAt:        v.UpdatedAt.Format(time_parse.CSTLayout),
				UpdatedUser:      v.UpdatedUser,
			}

			res.List[k] = data
		}

		ctx.Payload(res)
	}
}
