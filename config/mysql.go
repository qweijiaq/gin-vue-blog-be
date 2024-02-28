package config

import "strconv"

type Mysql struct {
	Host     string `yaml:"host"`      // 服务器地址
	Port     int    `yaml:"port"`      // 端口
	Config   string `yaml:"config"`    // 高级配置，例如 charset
	DB       string `yaml:"db"`        // 数据库名
	User     string `yaml:"user"`      // 数据库用户名
	Password string `yaml:"password"`  // 数据库密码
	LogLevel string `yaml:"log_level"` // Gorm 日志的级别
}

// Dsn 给 Mysql 结构体绑定一个方法：该方法返回连接数据库时需要传入的 DSN (数据来源名称)
func (m Mysql) Dsn() string {
	return m.User + ":" + m.Password + "@tcp(" + m.Host + ":" + strconv.Itoa(m.Port) + ")/" + m.DB + "?" + m.Config
}
