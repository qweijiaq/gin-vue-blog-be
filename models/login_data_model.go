package models

import "server/models/ctype"

// LoginDataModel 统计用户登录数据 id, 用户id, 用户昵称，用户token，登录设备，登录时间
type LoginDataModel struct {
	MODEL
	UserID    uint             `gorm:"comment:用户id" json:"user_id"`
	UserModel UserModel        `gorm:"foreignKey:UserID" json:"-"`
	IP        string           `gorm:"size:20;comment:ip" json:"ip"` // 登录的ip
	NickName  string           `gorm:"size:42;comment:昵称" json:"nick_name"`
	Token     string           `gorm:"size:256;comment:token" json:"token"`
	Device    string           `gorm:"size:256;comment:登录失败" json:"device"` // 登录设备
	Addr      string           `gorm:"size:64;comment:地址" json:"addr"`
	LoginType ctype.SignStatus `gorm:"size:type=smallint(6);comment:登录方式，1QQ，3邮箱" json:"login_type"` // 登录方式
}
