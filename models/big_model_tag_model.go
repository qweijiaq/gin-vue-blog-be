package models

// BigModelTagModel 大模型标签表
type BigModelTagModel struct {
	MODEL
	Title string              `gorm:"size:16" json:"title"` // 标签名称
	Color string              `gorm:"size:16" json:"color"` // 标签颜色
	Rules []BigModelRuleModel `gorm:"many2many:big_model_rule_tag_models" json:"roles"`
}
