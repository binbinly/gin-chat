package model

// EmoticonModel 表情
type EmoticonModel struct {
	PriID
	Category string `gorm:"column:category;type:varchar(100);not null;comment:分类" json:"category"`
	Name     string `gorm:"column:name;type:varchar(100);not null;comment:名称" json:"name"`
	Url      string `gorm:"column:url;type:varchar(256);not null;comment:路径" json:"url"`
	CUT
}

// Emoticon 表情包对外结构
type Emoticon struct {
	ID       int    `json:"id"`
	Category string `json:"category"`
	Name     string `json:"name"`
	Url      string `json:"url"`
}

// TableName 表名
func (f *EmoticonModel) TableName() string {
	return "emoticon"
}
