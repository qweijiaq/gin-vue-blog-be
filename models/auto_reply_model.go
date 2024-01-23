package models

// AutoReplyModel 自动回复表
type AutoReplyModel struct {
	MODEL
	Name         string `gorm:"size:32" json:"name"`            // 规则名称
	Mode         int    `json:"mode"`                           // 匹配模式 1-精确匹配 2-模糊匹配 3-前缀匹配 4-正则匹配
	Rule         string `gorm:"size:64" json:"rule"`            // 匹配规则
	ReplyContent string `gorm:"size:1024" json:"reply_content"` // 回复内容
}
