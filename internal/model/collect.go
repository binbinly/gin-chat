package model

const (
	// CollTypeEmo 表情
	CollTypeEmo = iota + 1
	// CollTypeText 文本
	CollTypeText
	// CollTypeImage 图片
	CollTypeImage
	// CollTypeVideo 视频
	CollTypeVideo
	// CollTYpeAudio 语音
	CollTYpeAudio
	// CollTypeCard 名片
	CollTypeCard
)

// CollectModel 用户收藏
type CollectModel struct {
	PriID
	UID
	Content string `gorm:"column:content;not null;size:5000;comment:内容" json:"content"`
	Type    int8   `gorm:"column:type;not null;comment:类型" json:"type"`
	Options string `gorm:"column:options;size:255;not null;default:'';comment:选项" json:"options"`
	CT
}

// TableName 表名
func (c *CollectModel) TableName() string {
	return "collect"
}
