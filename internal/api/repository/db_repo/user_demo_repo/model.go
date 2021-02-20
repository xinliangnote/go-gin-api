package user_demo_repo

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
