package model

const (
	// MomentSeeTypeAll 公开
	MomentSeeTypeAll = iota + 1
	// MomentSeeTypeNone 私密
	MomentSeeTypeNone
	// MomentSeeTypeOnly 指定可看
	MomentSeeTypeOnly
	// MomentSeeTypeExcept 不给谁看
	MomentSeeTypeExcept
)

const (
	// MomentTypeText 文本
	MomentTypeText = iota + 1
	// MomentTypeImage 图文
	MomentTypeImage
	// MomentTypeVideo 视频
	MomentTypeVideo
)

// MomentModel 朋友圈模型
type MomentModel struct {
	PriID
	UID
	Content  string `gorm:"column:content;not null;size:1000;comment:内容" json:"content"`
	Image    string `gorm:"column:image;not null;size:1000;default:'';comment:图片" json:"image"`
	Video    string `gorm:"column:video;not null;size:255;default:'';comment:视频地址" json:"video"`
	Location string `gorm:"column:location;not null;size:255;default:'';comment:地址" json:"location"`
	Remind   string `gorm:"column:remind;not null;size:255;default:'';comment:提醒谁看" json:"remind"`
	See      string `gorm:"column:see;not null;size:255;default:'';comment:用户id列表" json:"see"`
	Type     int8   `gorm:"column:type;not null;default:1;comment:动态类型" json:"type"`
	SeeType  int8   `gorm:"column:see_type;not null;default:1;comment:可见类型" json:"see_type"`
	CUT
}

// TableName 表名
func (m *MomentModel) TableName() string {
	return "moment"
}

// Moment 对外朋友圈结构体
type Moment struct {
	User *User         `json:"user"`
	List []*MomentList `json:"list"`
}

// MomentList 朋友圈列表结构体
type MomentList struct {
	ID        int        `json:"id"`
	Content   string     `json:"content"`
	Image     string     `json:"image"`
	Video     string     `json:"video"`
	Location  string     `json:"location"`
	Type      int8       `json:"type"`
	Likes     []*User    `json:"likes"`
	Comments  []*Comment `json:"comments"`
	User      *User      `json:"user"`
	CreatedAt int64      `json:"created_at"`
}

// Comment 朋友圈评论结构
type Comment struct {
	Content string `json:"content"`
	User    *User  `json:"user"`
	Reply   *User  `json:"reply"`
}
