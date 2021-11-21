package admin

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/xinliangnote/go-gin-api/configs"
	"github.com/xinliangnote/go-gin-api/internal/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/password"
	"github.com/xinliangnote/go-gin-api/internal/repository/redis"
	"github.com/xinliangnote/go-gin-api/internal/services/admin"
	"github.com/xinliangnote/go-gin-api/pkg/errno"
	"github.com/xinliangnote/go-gin-api/pkg/errors"
)

type loginRequest struct {
	Username string `form:"username"` // 用户名
	Password string `form:"password"` // 密码
}

type loginResponse struct {
	Token string `json:"token"` // 用户身份标识
}

// Login 管理员登录
// @Summary 管理员登录
// @Description 管理员登录
// @Tags API.admin
// @Accept multipart/form-data
// @Produce json
// @Param username formData string true "用户名"
// @Param password formData string true "密码"
// @Success 200 {object} loginResponse
// @Failure 400 {object} code.Failure
// @Router /api/admin/login [post]
func (h *handler) Login() core.HandlerFunc {
	return func(c core.Context) {
		req := new(loginRequest)
		res := new(loginResponse)
		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithErr(err),
			)
			return
		}

		searchOneData := new(admin.SearchOneData)
		searchOneData.Username = req.Username
		searchOneData.Password = password.GeneratePassword(req.Password)
		searchOneData.IsUsed = 1

		info, err := h.adminService.Detail(c, searchOneData)
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.AdminLoginError,
				code.Text(code.AdminLoginError)).WithErr(err),
			)
			return
		}

		if info == nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.AdminLoginError,
				code.Text(code.AdminLoginError)).WithErr(errors.New("未查询出符合条件的用户")),
			)
			return
		}

		token := password.GenerateLoginToken(info.Id)

		adminCacheData := &struct {
			Id       int32  `json:"id"`       // 主键ID
			Username string `json:"username"` // 用户名
			Nickname string `json:"nickname"` // 昵称
			Mobile   string `json:"mobile"`   // 手机号
		}{
			Id:       info.Id,
			Username: info.Username,
			Nickname: info.Nickname,
			Mobile:   info.Mobile,
		}

		// 用户信息
		adminJsonInfo, _ := json.Marshal(adminCacheData)

		// 将用户信息记录到 Redis 中
		err = h.cache.Set(configs.RedisKeyPrefixLoginUser+token, string(adminJsonInfo), time.Hour*24, redis.WithTrace(c.Trace()))
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.AdminLoginError,
				code.Text(code.AdminLoginError)).WithErr(err),
			)
			return
		}

		searchMenuData := new(admin.SearchMyMenuData)
		searchMenuData.AdminId = info.Id
		menu, err := h.adminService.MyMenu(c, searchMenuData)
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.AdminLoginError,
				code.Text(code.AdminLoginError)).WithErr(err),
			)
			return
		}

		// 菜单栏信息
		menuJsonInfo, _ := json.Marshal(menu)

		// 将菜单栏信息记录到 Redis 中
		err = h.cache.Set(configs.RedisKeyPrefixLoginUser+token+":menu", string(menuJsonInfo), time.Hour*24, redis.WithTrace(c.Trace()))
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.AdminLoginError,
				code.Text(code.AdminLoginError)).WithErr(err),
			)
			return
		}

		searchActionData := new(admin.SearchMyActionData)
		searchActionData.AdminId = info.Id
		action, err := h.adminService.MyAction(c, searchActionData)
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.AdminLoginError,
				code.Text(code.AdminLoginError)).WithErr(err),
			)
			return
		}

		// 可访问接口信息
		actionJsonInfo, _ := json.Marshal(action)

		// 将可访问接口信息记录到 Redis 中
		err = h.cache.Set(configs.RedisKeyPrefixLoginUser+token+":action", string(actionJsonInfo), time.Hour*24, redis.WithTrace(c.Trace()))
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.AdminLoginError,
				code.Text(code.AdminLoginError)).WithErr(err),
			)
			return
		}

		res.Token = token
		c.Payload(res)
	}
}
