package tool

import (
	"encoding/json"
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/validation"
	"github.com/xinliangnote/go-gin-api/internal/websocket/sysmessage"
	"github.com/xinliangnote/go-gin-api/pkg/errno"
	"github.com/xinliangnote/go-gin-api/pkg/timeutil"
)

type sendMessageRequest struct {
	Message string `form:"message"` // 消息内容
}

type sendMessageResponse struct {
	Status string `json:"status"` // 状态
}

// SendMessage 发送消息
// @Summary 发送消息
// @Description 发送消息
// @Tags API.tool
// @Accept multipart/form-data
// @Produce json
// @Param message formData string true "消息内容"
// @Success 200 {object} sendMessageResponse
// @Failure 400 {object} code.Failure
// @Router /api/tool/send_message [post]
func (h *handler) SendMessage() core.HandlerFunc {
	type messageBody struct {
		Username string `json:"username"`
		Message  string `json:"message"`
		Time     string `json:"time"`
	}

	return func(ctx core.Context) {
		req := new(sendMessageRequest)
		res := new(sendMessageResponse)
		if err := ctx.ShouldBindForm(req); err != nil {
			ctx.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ParamBindError,
				validation.Error(err)).WithErr(err),
			)
			return
		}

		conn, err := sysmessage.GetConn()
		if err != nil {
			ctx.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.SocketConnectError,
				code.Text(code.SocketConnectError)).WithErr(err),
			)
			return
		}

		messageData := new(messageBody)
		messageData.Username = ctx.UserName()
		messageData.Message = req.Message
		messageData.Time = timeutil.CSTLayoutString()

		messageJsonData, err := json.Marshal(messageData)
		if err != nil {
			ctx.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.SocketSendError,
				code.Text(code.SocketSendError)).WithErr(err),
			)
			return
		}

		err = conn.OnSend(messageJsonData)
		if err != nil {
			ctx.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.SocketSendError,
				code.Text(code.SocketSendError)).WithErr(err),
			)
			return
		}

		res.Status = "OK"
		ctx.Payload(res)
	}
}
