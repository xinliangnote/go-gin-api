package code

// 错误时返回结构
type Failure struct {
	Code    int    `json:"code"`    // 业务码
	Message string `json:"message"` // 描述信息
}

const (
	// 服务级错误码
	ServerError        = 10101
	TooManyRequests    = 10102
	ParamBindError     = 10103
	AuthorizationError = 10104
	CallHTTPError      = 10105
	ResubmitError      = 10106
	ResubmitMsg        = 10107
	HashIdsDecodeError = 10108

	// 模块级错误码 - 用户模块
	IllegalUserName = 20101
	UserCreateError = 20102
	UserUpdateError = 20103
	UserSearchError = 20104

	// 授权调用方
	AuthorizedCreateError    = 30101
	AuthorizedListError      = 30102
	AuthorizedDeleteError    = 30103
	AuthorizedUpdateError    = 30104
	AuthorizedDetailError    = 30105
	AuthorizedCreateAPIError = 30106
	AuthorizedListAPIError   = 30107
	AuthorizedDeleteAPIError = 30108
)

var codeText = map[int]string{
	ServerError:        "Internal Server Error",
	TooManyRequests:    "Too Many Requests",
	ParamBindError:     "参数信息有误",
	AuthorizationError: "签名信息有误",
	CallHTTPError:      "调用第三方 HTTP 接口失败",
	ResubmitError:      "Resubmit Error",
	ResubmitMsg:        "请勿重复提交",
	HashIdsDecodeError: "ID 参数有误",

	IllegalUserName: "非法用户名",
	UserCreateError: "创建用户失败",
	UserUpdateError: "更新用户失败",
	UserSearchError: "查询用户失败",

	AuthorizedCreateError:    "创建调用方失败",
	AuthorizedListError:      "获取调用方列表页失败",
	AuthorizedDeleteError:    "删除调用方失败",
	AuthorizedUpdateError:    "更新调用方失败",
	AuthorizedDetailError:    "获取调用方详情失败",
	AuthorizedCreateAPIError: "创建调用方API地址失败",
	AuthorizedListAPIError:   "获取调用方API地址列表失败",
	AuthorizedDeleteAPIError: "删除调用方API地址失败",
}

func Text(code int) string {
	return codeText[code]
}
