package config

import "fmt"

type Redis struct {
	IP       string `json:"ip" yaml:"ip"`               // IP
	Port     int    `json:"port" yaml:"port"`           // 端口
	Password string `json:"password" yaml:"password"`   // 密码
	PoolSize int    `json:"pool_size" yaml:"pool_size"` // 连接池大小
}

// Addr 返回 Redis 数据库的地址
func (r Redis) Addr() string {
	return fmt.Sprintf("%s:%d", r.IP, r.Port)
}
