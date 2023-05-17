package model

// FriendModel 好友模型
type FriendModel struct {
	PriID
	UID
	FriendID int    `gorm:"column:friend_id;type:int(11) unsigned;not null;uniqueIndex:idx_uid_fid;comment:好友id" json:"friend_id"`
	Nickname string `gorm:"column:nickname;type:varchar(60);not null;comment:备注昵称" json:"nickname"`
	LookMe   int8   `gorm:"column:look_me;not null;default:1;comment:看我" json:"look_me"`
	LookHim  int8   `gorm:"column:look_him;not null;default:1;comment:看他" json:"look_him"`
	IsStar   int8   `gorm:"column:is_star;not null;default:0;comment:是否星标用户" json:"is_star"`
	IsBlack  int8   `gorm:"column:is_black;not null;default:0;comment:是否拉黑" json:"is_black"`
	Tags     string `gorm:"column:tags;type:varchar(1000);not null;default:'';comment:标签" json:"tags"`
	CUT
	//Friend *UserModel `json:"friend" gorm:"foreignkey:id;references:friend_id"`
}

// TableName 表名
func (f *FriendModel) TableName() string {
	return "friend"
}

// Friend 对外暴露的好友信息结构体
type Friend struct {
	LookMe   int8     `json:"look_me"`
	LookHim  int8     `json:"look_him"`
	IsStar   int8     `json:"is_star"`
	IsBlack  int8     `json:"is_black"`
	Nickname string   `json:"nickname"`
	Tags     []string `json:"tags"`
}
