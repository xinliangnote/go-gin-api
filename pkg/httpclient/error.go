package httpclient

var _ ReplyErr = (*replyErr)(nil)

// ReplyErr 错误响应，当resp.StatusCode != http.StatusOK时用来包装返回的httpcode和body。
type ReplyErr interface {
	error
	StatusCode() int
	Body() []byte
}

type replyErr struct {
	err        error
	statusCode int
	body       []byte
}

func (r *replyErr) Error() string {
	return r.err.Error()
}

func (r *replyErr) StatusCode() int {
	return r.statusCode
}

func (r *replyErr) Body() []byte {
	return r.body
}

func newReplyErr(statusCode int, body []byte, err error) ReplyErr {
	return &replyErr{
		statusCode: statusCode,
		body:       body,
		err:        err,
	}
}

// ToReplyErr 尝试将err转换为ReplyErr
func ToReplyErr(err error) (ReplyErr, bool) {
	if err == nil {
		return nil, false
	}

	e, ok := err.(ReplyErr)
	return e, ok
}
