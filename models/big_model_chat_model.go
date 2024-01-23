package models

// BigModelChatModel 大模型对话表
type BigModelChatModel struct {
	MODEL
	SessionID    uint                 `json:"sessionID"`                     // 会话 ID
	SessionModel BigModelSessionModel `gorm:"foreignKey:SessionID" json:"-"` // 会话
	Status       bool                 `json:"status"`                        // 状态，AI 有没有正常的回复用户
	Content      string               `json:"content"`                       // 用户的聊天内容
	BotContent   string               `json:"botContent"`                    // AI 的回复内容
	RoleID       uint                 `json:"roleID"`                        // 是哪一个角色
	RoleModel    BigModelRoleModel    `gorm:"foreignKey:RoleID" json:"-"`    // 角色
	UserID       uint                 `json:"userID"`                        // 用户 ID
	UserModel    UserModel            `gorm:"foreignKey:UserID" json:"-"`    // 用户
}
