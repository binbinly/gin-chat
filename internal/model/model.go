package model

import "gorm.io/gorm"

const (
	// ReleaseYes 已发布
	ReleaseYes = 1
	// DefaultOrder 默认排序
	DefaultOrder = "id DESC"
)

const (
	// StatusInit 状态-初始化
	StatusInit = iota
	// StatusSuccess 状态-成功
	StatusSuccess
	// StatusError 状态-失败
	StatusError
)

// CUT 公共时间字段
type CUT struct {
	CreatedAt int64 `gorm:"column:created_at;not null;autoCreateTime;comment:创建时间" json:"created_at"`
	UpdatedAt int64 `gorm:"column:updated_at;not null;autoUpdateTime;comment:更新时间" json:"updated_at"`
}

// PriID 主键
type PriID struct {
	ID int `gorm:"primaryKey;autoIncrement;type:int;not null;column:id;comment:ID" json:"id"`
}

// CT 创建时间
type CT struct {
	CreatedAt int64 `gorm:"column:created_at;not null;autoCreateTime;comment:创建时间" json:"created_at"`
}

// UT 更新时间
type UT struct {
	UpdatedAt int64 `gorm:"column:updated_at;not null;autoUpdateTime;comment:更新时间" json:"updated_at"`
}

// DT 删除时间
type DT struct {
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间" json:"deleted_at"`
}

// UID 用户ID
type UID struct {
	UserID int `gorm:"column:user_id;not null;type:int;index;comment:用户id" json:"user_id"`
}

// OffsetPage 分页查询
func OffsetPage(offset, limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(limit)
	}
}
