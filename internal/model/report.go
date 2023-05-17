package model

const (
	// ReportStatusPending 待处理
	ReportStatusPending = iota + 1
	// ReportStatusFinish 已完成
	ReportStatusFinish
)

// ReportModel 用户举报模型
type ReportModel struct {
	PriID
	UID
	TargetID   int    `gorm:"column:target_id;not null;type:int(11) unsigned;index;comment:目标id" json:"target_id"`
	TargetType int8   `gorm:"column:target_type;not null;default:1;comment:目标类型，1=用户，2=群组" json:"target_type"`
	Content    string `gorm:"column:content;not null;type:varchar(5000);comment:内容" json:"content"`
	Category   string `gorm:"column:category;not null;type:varchar(255);default:'';comment:分类" json:"category"`
	Status     int8   `gorm:"column:status;not null;default:1;comment:状态" json:"status"`
	CUT
}

// TableName 表名
func (u *ReportModel) TableName() string {
	return "report"
}
