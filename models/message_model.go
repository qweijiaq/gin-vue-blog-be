package models

// MessageModel 记录消息
type MessageModel struct {
	MODEL
	SendUserID       uint      `gorm:"primaryKey;comment:发送人ID" json:"send_user_id"` // 发送人id
	SendUserModel    UserModel `gorm:"foreignKey:SendUserID" json:"-"`
	SendUserNickName string    `gorm:"size:42;comment:发送人昵称" json:"send_user_nick_name"`
	SendUserAvatar   string    `gorm:"comment:发送人头像" json:"send_user_avatar"`

	RevUserID       uint      `gorm:"primaryKey;comment:接收人ID" json:"rev_user_id"` // 接收人id
	RevUserModel    UserModel `gorm:"foreignKey:RevUserID" json:"-"`
	RevUserNickName string    `gorm:"size:42;comment:接收人昵称" json:"rev_user_nick_name"`
	RevUserAvatar   string    `gorm:"comment:接收人头像" json:"rev_user_avatar"`
	IsRead          bool      `gorm:"default:false;comment:接收人是否查看" json:"is_read"` // 接收方是否查看
	Content         string    `gorm:"comment:消息内容" json:"content"`                  // 消息内容
}
