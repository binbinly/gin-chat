package model

const (
	// MessageChatTypeUser 用户聊天类型
	MessageChatTypeUser = iota + 1
	// MessageChatTypeGroup 群组聊天类型
	MessageChatTypeGroup
)

const (
	// MessageTypeSystem 系统消息
	MessageTypeSystem = iota + 1
	// MessageTypeText 文本
	MessageTypeText
	// MessageTypeImage 图片
	MessageTypeImage
	// MessageTypeVideo 视频
	MessageTypeVideo
	// MessageTYpeAudio 音频
	MessageTYpeAudio
	// MessageTypeEmoticon 表情
	MessageTypeEmoticon
	// MessageTypeCard 名片
	MessageTypeCard
)

// MessageModel 消息模型
type MessageModel struct {
	PriID
	UID
	ToID     int    `gorm:"column:to_id;not null;type:int(11) unsigned;index;comment:发送者" json:"to_id"`
	ChatType int8   `gorm:"column:chat_type;not null;default:1;comment:目标类型，1=用户，2=群组" json:"chat_type"`
	Type     int8   `gorm:"column:type;not null;default:1;comment:消息类型" json:"type"`
	Content  string `gorm:"column:content;not null;type:varchar(5000);comment:内容" json:"content"`
	Options  string `gorm:"column:options;type:varchar(255);not null;default:'';comment:选项" json:"options"`
	CT
}

// TableName 表名
func (u *MessageModel) TableName() string {
	return "message"
}
