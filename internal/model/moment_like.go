package model

// MomentLikeModel 朋友圈点赞模型
type MomentLikeModel struct {
	PriID
	UserID   int `gorm:"column:user_id;not null;type:int(11) unsigned;uniqueIndex:idx_uid_mid;comment:用户id" json:"user_id"`
	MomentID int `gorm:"column:moment_id;not null;type:int(11) unsigned;uniqueIndex:idx_uid_mid;comment:动态id" json:"moment_id"`
	CUT
}

// TableName 表名
func (m *MomentLikeModel) TableName() string {
	return "moment_like"
}

// MomentLike 朋友圈点赞结构
type MomentLike struct {
	UserID   int `json:"user_id"`
	MomentID int `json:"moment_id"`
}
