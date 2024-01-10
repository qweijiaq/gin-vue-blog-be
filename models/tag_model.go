package models

// TagModel 标签表
type TagModel struct {
	MODEL
	Title string `gorm:"size:16;comment:标签名" json:"title"` // 标签的名称
}
