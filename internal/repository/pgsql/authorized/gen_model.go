package authorized

import "time"

// Authorized 已授权的调用方表
//
//go:generate gormgen -structs Authorized -input .
type Authorized struct {
	Id                int32     // 主键
	BusinessKey       string    // 调用方key
	BusinessSecret    string    // 调用方secret
	BusinessDeveloper string    // 调用方对接人
	Remark            string    // 备注
	IsUsed            int32     // 是否启用 1:是  -1:否
	IsDeleted         int32     // 是否删除 1:是  -1:否
	CreatedAt         time.Time `gorm:"time"` // 创建时间
	CreatedUser       string    // 创建人
	UpdatedAt         time.Time `gorm:"time"` // 更新时间
	UpdatedUser       string    // 更新人
}
