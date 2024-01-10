package routers

import (
	"server/api"
	"server/middleware"
)

func (router RouterGroup) CommentRouter() {
	app := api.ApiGroupApp.CommentApi
	router.POST("comments", middleware.JwtAuth(), app.CommentCreateView)
	router.GET("comments", app.CommentListAllView)                                       // 评论列表
	router.GET("comments/articles", middleware.JwtAdmin(), app.CommentByArticleListView) // 有评论的文章列表
	router.GET("comments/:id", app.CommentListView)                                      // 文章下的评论列表
	router.GET("comments/digg/:id", app.CommentDigg)                                     // 评论点赞
	router.DELETE("comments/:id", middleware.JwtAuth(), app.CommentRemoveView)
}
