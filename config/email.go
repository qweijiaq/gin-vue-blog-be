package config

type Email struct {
	Host             string `json:"host" yaml:"host"`                             // 主机
	Port             int    `json:"port" yaml:"port"`                             // 端口
	User             string `json:"user" yaml:"user"`                             // 发件人邮箱
	Password         string `json:"password" yaml:"password"`                     // 授权码
	DefaultFromEmail string `json:"default_from_email" yaml:"default_from_email"` // 默认的发件人名字
	UseSSL           bool   `json:"use_ssl" yaml:"use_ssl"`                       // 是否使用 SSL
	UserTls          bool   `json:"user_tls" yaml:"user_tls"`
}
