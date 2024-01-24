package routers

import (
	"server/api"
	"server/middleware"
)

func (router RouterGroup) BigModelRouter() {
	app := api.ApiGroupApp.BigModelApi

	// 配置相关
	{
		router.GET("bigModel/usable", middleware.JwtAdmin(), app.ModelUsableListView)                    // 获取可用大模型列表
		router.GET("bigModel/setting", middleware.JwtAdmin(), app.ModelSettingView)                      // 获取大模型配置
		router.PUT("bigModel/setting", middleware.JwtAdmin(), app.ModelSettingUpdateView)                // 更新大模型配置
		router.GET("bigModel/session_setting", middleware.JwtAdmin(), app.ModelSessionSettingView)       // 获取大模型会话配置
		router.PUT("bigModel/session_setting", middleware.JwtAdmin(), app.ModelSessionSettingUpdateView) // 更新大模型会话配置
		router.PUT("bigModel/auto_reply", middleware.JwtAdmin(), app.AutoReplyUpdateView)                // 自动回复规则添加与更新
		router.GET("bigModel/auto_reply", middleware.JwtAdmin(), app.AutoReplyListView)                  // 获取自动回复规则列表
		router.DELETE("bigModel/auto_reply", middleware.JwtAdmin(), app.AutoReplyRemoveView)             // 删除自动回复规则
	}

	// 用户相关
	{
		router.GET("bigModel/user_scope_enable", middleware.JwtAuth(), app.UserScopeEnableView) // 获取用户是否可以领取积分
		router.POST("bigModel/user_scope", middleware.JwtAuth(), app.UserScopeView)             // 用户领取积分

	}

	// 角色相关
	{
		router.GET("bigModel/roles/:id", middleware.JwtAuth(), app.RoleListView)               // 角色详情
		router.POST("bigModel/roles", middleware.JwtAdmin(), app.RoleCreateView)               // 角色添加
		router.PUT("bigModel/roles", middleware.JwtAdmin(), app.RoleUpdateView)                // 角色更新
		router.GET("bigModel/roles", middleware.JwtAuth(), app.RoleListView)                   // 角色列表                      // 获取角色列表
		router.DELETE("bigModel/roles", middleware.JwtAdmin(), app.RoleRemoveView)             // 角色删除
		router.GET("bigModel/square", app.RoleSquareView)                                      // 角色广场
		router.GET("bigModel/role_history", middleware.JwtAuth(), app.RoleUserHistoryListView) // 角色历史列表
		router.PUT("bigModel/tags", middleware.JwtAdmin(), app.TagUpdateView)                  // 角色标签添加与更新
		router.GET("bigModel/tags", middleware.JwtAdmin(), app.TagListView)                    // 获取角色标签列表
		router.DELETE("bigModel/tags", middleware.JwtAdmin(), app.TagRemoveView)               // 角色标签删除
	}

	// 会话相关
	{
		router.POST("bigModel/sessions", middleware.JwtAuth(), app.SessionCreateView)          // 用户创建会话
		router.GET("bigModel/sessions", middleware.JwtAuth(), app.SessionListView)             // 用户获取会话列表
		router.PUT("bigModel/sessions", middleware.JwtAuth(), app.SessionUpdateNameView)       // 用户修改会话名称
		router.DELETE("bigModel/session/:id", middleware.JwtAuth(), app.SessionUserRemoveView) // 用户删除会话
		router.DELETE("bigModel/sessions", middleware.JwtAdmin(), app.SessionRemoveView)       // 管理员删除会话
	}

	// 对话相关
	{
		router.GET("bigModel/chats", middleware.JwtAuth(), app.ChatListView)              // 对话列表
		router.POST("bigModel/chats", middleware.JwtAuth(), app.ChatCreateView)           // 用户创建对话
		router.DELETE("bigModel/chats/:id", middleware.JwtAuth(), app.ChatUserRemoveView) // 用户删除对话
	}
}
