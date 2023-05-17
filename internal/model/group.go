package model

// GroupModel 群组模型
type GroupModel struct {
	PriID
	UID
	Name          string `gorm:"column:name;type:varchar(255);not null;comment:群组名" json:"name"`
	Avatar        string `gorm:"column:avatar;not null;type:varchar(128);default:'';comment:头像" json:"avatar"`
	Remark        string `gorm:"column:remark;not null;default:'';type:varchar(500);comment:备注" json:"remark"`
	InviteConfirm int8   `gorm:"column:invite_confirm;not null;default:0;comment:邀请确认" json:"invite_confirm"`
	Status        int8   `gorm:"column:status;not null;default:1;comment:状态" json:"status"`
	CUT
	DT
}

// TableName 表名
func (g *GroupModel) TableName() string {
	return "group"
}

// GroupList 群列表
type GroupList struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}
