package main

import (
	"server/core"
	_ "server/docs"
	"server/flags"
	"server/global"
	"server/routers"
	"server/service/cron"
	"server/utils"
)

// @title gvb_server API 文档
// @version 1.0
// @description gvb_server API 文档
// @host 127.0.0.1:3001
// @BasePath /
func main() {
	// 读取配置文件
	core.InitConf()
	// 初始化日志
	global.Log = core.InitLogger()
	// 连接数据库
	global.DB = core.InitGorm()

	core.InitAddrDB()

	// 连接 Redis
	global.Redis = core.ConnectRedis()
	// 连接 ES
	global.ESClient = core.EsConnect()

	defer global.AddrDB.Close()

	// 命令行参数绑定
	option := flags.Parse()
	if option.Run() {
		return
	}

	cron.CronInit()

	router := routers.InitRouter()

	addr := global.Config.System.Addr()

	utils.PrintSystem()

	err := router.Run(addr)
	if err != nil {
		global.Log.Fatalf(err.Error())
	}
}
