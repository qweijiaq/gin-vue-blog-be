package models

// AdvertModel 广告表
type AdvertModel struct {
	MODEL
	Title  string `gorm:"size:32;comment:广告标题" json:"title"` // 显示的标题
	Href   string `gorm:"comment:跳转链接" json:"href"`          // 跳转链接
	Images string `gorm:"comment:图片" json:"images"`          // 图片
	IsShow bool   `gorm:"comment:是否展示" json:"is_show"`       // 是否展示
}
