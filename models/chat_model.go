package models

import "server/models/ctype"

type ChatModel struct {
	MODEL    `json:","`
	NickName string        `gorm:"size:15;comment:昵称" json:"nick_name"`
	Avatar   string        `gorm:"size:128;comment:头像" json:"avatar"`
	Content  string        `gorm:"size:256;comment:内容" json:"content"`
	IP       string        `gorm:"size:32;comment:ip" json:"ip,omit(list)"`
	Addr     string        `gorm:"size:64;comment:地址" json:"addr,omit(list)"`
	IsGroup  bool          `gorm:"comment:是否是群组消息" json:"is_group"` // 是否是群组消息
	MsgType  ctype.MsgType `gorm:"size:4;comment:消息类型" json:"msg_type"`
}
