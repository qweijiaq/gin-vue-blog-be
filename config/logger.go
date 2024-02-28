package config

type Logger struct {
	Level        string `yaml:"level"`          // 日志等级
	Prefix       string `yaml:"prefix"`         // 前缀
	Director     string `yaml:"director"`       // 保存日志的文件夹名
	ShowLine     bool   `yaml:"show_line"`      // 是否显示行号
	LogInConsole bool   `yaml:"log_in_console"` // 是否显示打印的路径
}
