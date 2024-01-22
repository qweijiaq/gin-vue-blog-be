package models

// AutoReplyModel 自动回复表
type AutoReplyModel struct {
	MODEL
	Name         string `gorm:"size:32" json:"name"`
	Mode         int    `json:"mode"`
	Rule         string `gorm:"size:64" json:"rule"`
	ReplyContent string `gorm:"size:1024", json:"reply_content"`
}
