package routers

import (
	"server/api"
	"server/middleware"
)

func (router RouterGroup) MessageRouter() {
	app := api.ApiGroupApp.MessageApi
	router.POST("messages", middleware.JwtAuth(), app.MessageCreateView)
	router.GET("messages_all", app.MessageListAllView)
	router.GET("messages", middleware.JwtAuth(), app.MessageListView)
	router.POST("messages_record", middleware.JwtAuth(), app.MessageRecordView)
	router.GET("message_users", middleware.JwtAdmin(), app.MessageUserListView)
	router.DELETE("message_users", middleware.JwtAuth(), app.MessageRecordRemoveView) // 删除聊天记录
	router.GET("message_users/user", middleware.JwtAdmin(), app.MessageUserListByUserView)
	router.GET("message_users/record", middleware.JwtAdmin(), app.MessageUserRecordView)
	router.GET("message_users/me", middleware.JwtAuth(), app.MessageUserListByMeView)
	router.GET("message_users/record/me", middleware.JwtAuth(), app.MessageUserRecordByMeView)
}
