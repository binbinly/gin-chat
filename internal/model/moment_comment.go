package model

// MomentCommentModel 朋友圈评论模型
type MomentCommentModel struct {
	PriID
	UID
	ReplyID  int    `gorm:"column:reply_id;not null;type:int;index;comment:回复用户id" json:"reply_id"`
	MomentID int    `gorm:"column:moment_id;not null;type:int;index;comment:动态id" json:"moment_id"`
	Content  string `gorm:"column:content;not null;size:1000;comment:评论内容" json:"content"`
	CUT
}

// TableName 表名
func (m *MomentCommentModel) TableName() string {
	return "moment_comment"
}
