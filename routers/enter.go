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
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()

	router.Use(middleware.LogMiddleWare())
	router.StaticFS("uploads", http.Dir("uploads"))
	router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	apiRouterGroup := router.Group("api")

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

	return router
}
