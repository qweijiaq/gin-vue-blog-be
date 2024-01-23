package models

// BigModelRoleModel 大模型角色表
type BigModelRoleModel struct {
	MODEL
	Name      string             `gorm:"size:16" json:"name"`                             // 角色名称
	Enable    bool               `json:"enable"`                                          // 是否启用
	Icon      string             `gorm:"size:256" json:"icon"`                            // 可以选择一些默认的，也可以图片上传
	Abstract  string             `gorm:"size:512" json:"abstract"`                        // 角色简介
	Tags      []BigModelTagModel `gorm:"many2many:big_model_role_tag_models" json:"tags"` // 标签列表
	Scope     int                `json:"scope"`                                           // 消耗的积分
	Prologue  string             `gorm:"size:512" json:"prologue"`                        // 开场白
	Prompt    string             `gorm:"size:2048" json:"prompt"`                         // 设定词
	AutoReply bool               `json:"auto_reply"`                                      // 是否接入自动回复
}
