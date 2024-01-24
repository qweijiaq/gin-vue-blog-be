package models

// BigModelSessionModel 大模型会话表
type BigModelSessionModel struct {
	MODEL
	Name      string              `json:"name"`   // 会话名称
	UserID    uint                `json:"userID"` // 用户 ID
	UserModel UserModel           `gorm:"foreignKey:UserID" json:"-"`
	RoleID    uint                `json:"roleID"` // 角色 ID
	RoleModel BigModelRoleModel   `gorm:"foreignKey:RoleID" json:"-"`
	ChatList  []BigModelChatModel `gorm:"foreignKey:SessionID" json:"-"` // 会话列表
}
