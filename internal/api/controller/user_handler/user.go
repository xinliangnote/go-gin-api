package user_handler

import (
	"errors"
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/api/code"
	"github.com/xinliangnote/go-gin-api/internal/api/model/user_model"
	"github.com/xinliangnote/go-gin-api/internal/api/service/user_service"
	"github.com/xinliangnote/go-gin-api/internal/pkg/cache"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/db"
	"github.com/xinliangnote/go-gin-api/pkg/ddm"
	"github.com/xinliangnote/go-gin-api/pkg/errno"

	"go.uber.org/zap"
)

var _ UserDemo = (*userDemo)(nil)

type UserDemo interface {
	// i 为了避免被其他包实现
	i()

	// Create 创建用户
	Create() core.HandlerFunc

	// UpdateNickNameByID 编辑用户 - 通过主键ID更新用户昵称
	UpdateNickNameByID() core.HandlerFunc

	// Delete 删除用户 - 通过主键ID更新 is_deleted = 1
	Delete() core.HandlerFunc

	// Detail 用户详情
	Detail() core.HandlerFunc
}

type userDemo struct {
	logger      *zap.Logger
	cache       cache.Repo
	userService user_service.UserService
}

func NewUserDemo(logger *zap.Logger, db db.Repo, cache cache.Repo) UserDemo {
	return &userDemo{
		logger:      logger,
		cache:       cache,
		userService: user_service.NewUserService(db, cache),
	}
}

func (u *userDemo) i() {}

// 创建用户
// @Summary 创建用户
// @Description 创建用户
// @Tags User
// @Accept  json
// @Produce  json
// @Param RequestInfo body user_model.CreateRequest true "请求信息"
// @Param Authorization header string true "签名"
// @Success 200 {object} user_model.CreateResponse
// @Failure 400 {object} code.Failure
// @Failure 401 {object} code.Failure
// @Router /user/create [post]
func (u *userDemo) Create() core.HandlerFunc {
	return func(c core.Context) {
		req := new(user_model.CreateRequest)
		res := new(user_model.CreateResponse)
		if err := c.ShouldBindJSON(req); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithErr(err),
			)
			return
		}

		if req.UserName == "" {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.IllegalUserName,
				code.Text(code.IllegalUserName)).WithErr(errors.New("req.UserName = ''")),
			)
			return
		}

		id, err := u.userService.Create(c, req)
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.UserCreateError,
				code.Text(code.UserCreateError)).WithErr(err),
			)
			return
		}

		res.Id = id
		c.Payload(res)

	}
}

// 编辑用户 - 通过用户主键ID更新用户昵称
// @Summary 编辑用户 - 通过用户主键ID更新用户昵称
// @Description 编辑用户 - 通过用户主键ID更新用户昵称
// @Tags User
// @Accept  json
// @Produce  json
// @Param RequestInfo body user_model.UpdateNickNameByIDRequest true "请求信息"
// @Param Authorization header string true "签名"
// @Success 200 {object} user_model.UpdateNickNameByIDResponse
// @Failure 400 {object} code.Failure
// @Failure 401 {object} code.Failure
// @Router /user/update [put]
func (u *userDemo) UpdateNickNameByID() core.HandlerFunc {
	return func(c core.Context) {
		req := new(user_model.UpdateNickNameByIDRequest)
		res := new(user_model.UpdateNickNameByIDResponse)
		if err := c.ShouldBindJSON(req); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithErr(err),
			)
			return
		}

		err := u.userService.UpdateNickNameByID(c, req.Id, req.NickName)
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.UserUpdateError,
				code.Text(code.UserUpdateError)).WithErr(err),
			)
			return
		}

		res.Id = req.Id

		c.Payload(res)
	}
}

// 删除用户 - 更新 is_deleted = 1
// @Summary 删除用户 - 更新 is_deleted = 1
// @Description 删除用户 - 更新 is_deleted = 1
// @Tags User
// @Accept  json
// @Produce  json
// @Param id path int true "用户ID"
// @Param Authorization header string true "签名"
// @Success 200
// @Failure 400 {object} code.Failure
// @Failure 401 {object} code.Failure
// @Router /user/delete/{id} [patch]
func (u *userDemo) Delete() core.HandlerFunc {
	return func(c core.Context) {
		req := new(user_model.DeleteRequest)
		if err := c.ShouldBindURI(req); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithErr(err),
			)
			return
		}

		err := u.userService.Delete(c, req.Id)
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.UserUpdateError,
				code.Text(code.UserUpdateError)).WithErr(err),
			)
			return
		}

		c.Payload("")
	}
}

// 用户详情
// @Summary 用户详情
// @Description 用户详情
// @Tags User
// @Accept  json
// @Produce  json
// @Param username path string true "用户名"
// @Param Authorization header string true "签名"
// @Success 200 {object} user_model.DetailResponse
// @Failure 400 {object} code.Failure
// @Failure 401 {object} code.Failure
// @Router /user/info/{username} [get]
func (u *userDemo) Detail() core.HandlerFunc {
	return func(c core.Context) {
		req := new(user_model.DetailRequest)
		res := new(user_model.DetailResponse)
		if err := c.ShouldBindURI(req); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithErr(err),
			)
			return
		}

		user, err := u.userService.GetUserByUserName(c, req.UserName)
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.UserSearchError,
				code.Text(code.UserSearchError)).WithErr(err),
			)
			return
		}

		res.Id = user.Id
		res.UserName = user.UserName
		res.NickName = user.NickName
		res.Mobile = ddm.Mobile(user.Mobile)

		c.Payload(res)
	}
}
