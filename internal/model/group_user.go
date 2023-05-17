package model

// GroupUserModel 群组用户模型
type GroupUserModel struct {
	PriID
	UID
	GroupID  int    `gorm:"column:group_id;type:int(11) unsigned;not null;comment:群组ID" json:"group_id"`
	Nickname string `gorm:"column:nickname;type:varchar(60);not null;comment:备注昵称" json:"nickname"`
	CUT
	DT
}

// TableName 表名
func (g *GroupUserModel) TableName() string {
	return "group_user"
}
