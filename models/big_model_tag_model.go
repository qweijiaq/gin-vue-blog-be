package models

// BigModelTagModel 大模型标签表
type BigModelTagModel struct {
	MODEL
	Title string              `gorm:"size:16" json:"title"` // 标签名称
	Color string              `gorm:"size:16" json:"color"` // 标签颜色
	Roles []BigModelRoleModel `gorm:"many2many:big_model_role_tag_models;joinForeignKey:big_model_tag_model_id;JoinReferences:big_model_role_model_id" json:"roles"`
}
