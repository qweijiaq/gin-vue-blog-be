package config

type Config struct {
	Mysql     Mysql     `yaml:"mysql"`      // MySQL 数据库配置
	Logger    Logger    `yaml:"logger"`     // 日志配置
	System    System    `yaml:"system"`     // 系统配置
	SiteInfo  SiteInfo  `yaml:"site_info"`  // 站点信息配置
	Upload    Upload    `yaml:"upload"`     // 图片上传配置
	QQ        QQ        `yaml:"qq"`         // QQ 第三方配置
	QiNiu     QiNiu     `yaml:"qi_niu"`     // 七牛云第三方配置
	Email     Email     `yaml:"email"`      // 邮箱配置
	Jwt       Jwt       `yaml:"jwt"`        // JWT 配置
	Redis     Redis     `yaml:"redis"`      // Redis 数据库配置
	ES        ES        `yaml:"es"`         // ElasticSearch 数据库配置
	ChatGroup ChatGroup `yaml:"chat_group"` // 群聊配置
	Gaode     Gaode     `yaml:"gaode"`      // 高德配置
	BigModel  BigModel  `yaml:"big_model"`  // 大模型配置
}
