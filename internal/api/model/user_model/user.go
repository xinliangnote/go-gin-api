package user_model

import (
	"time"
)

// 用户Demo表
type UserDemo struct {
	Id        uint      `gorm:"column:id;primary_key;AUTO_INCREMENT"`                 // 主键
	UserName  string    `gorm:"column:user_name;NOT NULL"`                            // 用户名
	NickName  string    `gorm:"column:nick_name;NOT NULL"`                            // 昵称
	Mobile    string    `gorm:"column:mobile;NOT NULL"`                               // 手机号
	IsDeleted int       `gorm:"column:is_deleted;default:-1;NOT NULL"`                // 是否删除 1:是  -1:否
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP;NOT NULL"` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;NOT NULL"` // 更新时间
}

func (m *UserDemo) TableName() string {
	return "user_demo"
}

// user_handler Create Request
type CreateRequest struct {
	UserName string `json:"user_name"` // 用户名
	NickName string `json:"nick_name"` // 昵称
	Mobile   string `json:"mobile"`    // 手机号
}

// user_handler Create Response
type CreateResponse struct {
	Id uint `json:"id"` // 主键ID
}

// user_handler UpdateNickNameByID Request
type UpdateNickNameByIDRequest struct {
	Id       uint   `json:"id"`        // 用户主键ID
	NickName string `json:"nick_name"` // 昵称
}

// user_handler UpdateNickNameByID Response
type UpdateNickNameByIDResponse struct {
	Id uint `json:"id"` // 用户主键ID
}

// user_handler Delete Request
type DeleteRequest struct {
	Id uint `uri:"id"` // 用户ID
}

// user_handler Detail Request
type DetailRequest struct {
	UserName string `uri:"username"` // 用户名
}

// user_handler Detail Response
type DetailResponse struct {
	Id       uint   `json:"id"`        // 用户主键ID
	UserName string `json:"user_name"` // 用户名
	NickName string `json:"nick_name"` // 昵称
	Mobile   string `json:"mobile"`    // 手机号
}
