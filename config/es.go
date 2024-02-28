package config

import (
	"fmt"
)

type ES struct {
	Host     string `yaml:"host"`     // 主机名
	Port     int    `yaml:"port"`     // 端口
	User     string `yaml:"user"`     // 用户名
	Password string `yaml:"password"` // 密码
}

// URL 返回 Redis 数据库的连接 URL
func (es ES) URL() string {
	return fmt.Sprintf("%s:%d", es.Host, es.Port)
}
