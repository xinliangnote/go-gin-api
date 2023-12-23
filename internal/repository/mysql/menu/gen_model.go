package menu

import "time"

// Menu 左侧菜单栏表
//
//go:generate gormgen -structs Menu -input .
type Menu struct {
	Id          int32     // 主键
	Pid         int32     // 父类ID
	Name        string    // 菜单名称
	Link        string    // 链接地址
	Icon        string    // 图标
	Level       int32     // 菜单类型 1:一级菜单 2:二级菜单
	Sort        int32     // 排序
	IsUsed      int32     // 是否启用 1:是 -1:否
	IsDeleted   int32     // 是否删除 1:是  -1:否
	CreatedAt   time.Time `gorm:"time"` // 创建时间
	CreatedUser string    // 创建人
	UpdatedAt   time.Time `gorm:"time"` // 更新时间
	UpdatedUser string    // 更新人
}
