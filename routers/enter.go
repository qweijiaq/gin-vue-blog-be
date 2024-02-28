package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"net/http"
	"server/global"
	"server/middleware"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env) // 设置项目的环境模式，可以时开发模式 dev 或线上模式 release
	router := gin.Default()

	router.Use(middleware.LogMiddleWare())
	// 保证后面上传的图片可以访问
	// 第一个参数是 web 的访问别名  第二个参数是内部的映射目录
	// 线上如果有 Nginx，这一步可以省略
	router.StaticFS("uploads", http.Dir("uploads"))

	router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	// 创建一个以 api 开头的路由分组 apiGroup -- 用于管理所有的路由
	apiRouterGroup := router.Group("api")

	// 又将这个路由分组赋给了 RouterGroup
	routerGroupApp := RouterGroup{apiRouterGroup}
	// 系统配置 API
	routerGroupApp.SettingsRouter()
	// 图片 API
	routerGroupApp.ImagesRouter()
	// 广告 API
	routerGroupApp.AdvertRouter()
	// 菜单 API
	routerGroupApp.MenuRouter()
	// 用户 API
	routerGroupApp.UserRouter()
	// 标签 API
	routerGroupApp.TagRouter()
	// 消息 API
	routerGroupApp.MessageRouter()
	// 文章 API
	routerGroupApp.ArticleRouter()
	// 评论 API
	routerGroupApp.CommentRouter()
	// 新闻 API
	routerGroupApp.NewsRouter()
	// 聊天 API
	routerGroupApp.ChatRouter()
	// 日志 API
	routerGroupApp.LogRouter()
	// 数据 API
	routerGroupApp.DataRouter()
	// 日志 API -- 2.0 升级版
	routerGroupApp.LogV2Router()
	// 角色 API
	routerGroupApp.RoleRouter()
	// 高德 API
	routerGroupApp.GaodeRouter()
	// 反馈 API
	routerGroupApp.FeedbackRouter()
	// 大模型 API
	routerGroupApp.BigModelRouter()

	return router
}
