package user_demo_repo

import "time"

// 用户Demo表
//go:generate gormgen -structs UserDemo -input .
type UserDemo struct {
	Id        int32     // 主键
	UserName  string    // 用户名
	NickName  string    // 昵称
	Mobile    string    // 手机号
	IsDeleted int32     // 是否删除 1:是  -1:否
	CreatedAt time.Time `gorm:"time"` // 创建时间
	UpdatedAt time.Time `gorm:"time"` // 更新时间
}
