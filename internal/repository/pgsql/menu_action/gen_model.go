package menu_action

import "time"

// MenuAction 功能权限表
//
//go:generate gormgen -structs MenuAction -input .
type MenuAction struct {
	Id          int32     // 主键
	MenuId      int32     // 菜单栏ID
	Method      string    // 请求方式
	Api         string    // 请求地址
	IsDeleted   int32     // 是否删除 1:是  -1:否
	CreatedAt   time.Time `gorm:"time"` // 创建时间
	CreatedUser string    // 创建人
	UpdatedAt   time.Time `gorm:"time"` // 更新时间
	UpdatedUser string    // 更新人
}
