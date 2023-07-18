package model

// UserTagModel 用户标签模型
type UserTagModel struct {
	PriID
	UserID int    `gorm:"column:user_id;not null;uniqueIndex:idx_name;type:int;index;comment:用户id" json:"user_id"`
	Name   string `gorm:"column:name;size:60;not null;uniqueIndex:idx_name;comment:标签名" json:"name"`
	CUT
}

// TableName 表名
func (g *UserTagModel) TableName() string {
	return "user_tag"
}

// UserTag 用户标签导出结构
type UserTag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
