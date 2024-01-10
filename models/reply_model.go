package models

type ReplyModel struct {
	MODEL
	Email        string `gorm:"size:64;comment:邮箱" json:"email"`
	Content      string `gorm:"size:128;comment:内容" json:"content"`
	ApplyContent string `gorm:"size:128;comment:回复内容" json:"apply_content"` // 回复的内容
	IsApply      bool   `gorm:"comment:是否回复" json:"is_apply"`               // 是否回复
}
