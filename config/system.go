package config

import "fmt"

type System struct {
	Host string `yaml:"host"` // 主机 IP
	Port int    `yaml:"port"` // 端口
	Env  string `yaml:"env"`  // 环境
}

func (s System) Addr() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}
