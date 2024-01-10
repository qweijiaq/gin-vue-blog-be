package flags

import (
	"flag"
)

type Option struct {
	DB     bool
	User   string // -u admin  -u user
	Load   string // 导入数据库文件
	Dump   bool   // 导出数据库
	Es     bool   // 创建索引
	ESDump bool   // 导出es索引
	ESLoad string // 导入es索引
}

// Parse 解析命令行参数
func Parse() (option *Option) {
	option = new(Option)
	flag.StringVar(&option.User, "u", "", "创建用户")
	flag.BoolVar(&option.DB, "db", false, "初始化数据库")
	flag.BoolVar(&option.Es, "es", false, "创建索引")
	flag.BoolVar(&option.Dump, "dump", false, "导出 SQL 数据库")
	flag.StringVar(&option.Load, "load", "", "导入 SQL 数据库")
	flag.BoolVar(&option.ESDump, "esdump", false, "导出 ES 索引")
	flag.StringVar(&option.ESLoad, "esload", "", "导入 ES 索引")
	flag.Parse()
	return option
}

// Run 根据命令执行不同的函数
func (option Option) Run() bool {
	if option.User != "" {
		CreateUser(option.User)
		return true
	}
	if option.DB {
		DB()
		return true
	}
	if option.Load != "" {
		Load(option.Load)
		return true
	}
	if option.Es {
		ESIndex()
		return true
	}
	if option.Dump {
		Dump()
		return true
	}
	if option.ESDump {
		ESDump()
		return true
	}
	if option.ESLoad != "" {
		ESLoad(option.ESLoad)
		return true
	}
	return false
}
