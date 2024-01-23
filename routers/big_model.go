package routers

import (
	"server/api"
	"server/middleware"
)

func (router RouterGroup) BigModelRouter() {
	app := api.ApiGroupApp.BigModelApi
	router.GET("bigModel/usable", middleware.JwtAdmin(), app.ModelUsableListView)                    // 获取可用大模型列表
	router.GET("bigModel/setting", app.ModelSettingView)                                             // 获取大模型配置
	router.PUT("bigModel/setting", middleware.JwtAdmin(), app.ModelSettingUpdateView)                // 更新大模型配置
	router.GET("bigModel/session_setting", middleware.JwtAdmin(), app.ModelSessionSettingView)       // 获取大模型会话配置
	router.PUT("bigModel/session_setting", middleware.JwtAdmin(), app.ModelSessionSettingUpdateView) // 更新大模型会话配置
	router.GET("bigModel/user_scope_enable", middleware.JwtAuth(), app.UserScopeEnableView)          // 获取用户是否可以领取积分
	router.POST("bigModel/user_scope", middleware.JwtAuth(), app.UserScopeView)                      // 用户领取积分
	router.PUT("bigModel/auto_reply", middleware.JwtAdmin(), app.AutoReplyUpdateView)                // 自动回复规则添加与更新
	router.GET("bigModel/auto_reply", middleware.JwtAdmin(), app.AutoReplyListView)                  // 获取自动回复规则列表
	router.DELETE("bigModel/auto_reply", middleware.JwtAdmin(), app.AutoReplyRemoveView)             // 删除自动回复规则

	router.PUT("bigModel/tags", middleware.JwtAdmin(), app.TagUpdateView) // 标签添加与更新
	router.GET("bigModel/tags", app.TagListView)                          // 获取标签列表
	router.DELETE("bigModel/tags", app.TagRemoveView)                     // 标签删除

	router.POST("bigModel/roles", middleware.JwtAdmin(), app.RoleCreateView) // 角色添加
	router.PUT("bigModel/roles", middleware.JwtAdmin(), app.RoleUpdateView)  // 角色更新
	router.GET("bigModel/roles", app.RoleListView)                           // 获取角色列表
}
