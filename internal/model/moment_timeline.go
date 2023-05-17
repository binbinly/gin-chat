package model

// MomentTimelineModel 朋友圈时间线模型
type MomentTimelineModel struct {
	PriID
	UID
	MomentID int  `gorm:"column:moment_id;not null;type:int(11) unsigned;index;comment:动态id" json:"moment_id"`
	IsOwn    int8 `gorm:"column:is_own;not null;default:0;comment:是否自己的" json:"is_own"`
	CUT
}

// TableName 表名
func (m *MomentTimelineModel) TableName() string {
	return "moment_timeline"
}
