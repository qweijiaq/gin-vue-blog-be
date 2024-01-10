package routers

import (
	"server/api"
	"server/middleware"
)

func (router RouterGroup) ArticleRouter() {
	app := api.ApiGroupApp.ArticleApi
	router.POST("articles", middleware.JwtAdmin(), app.ArticleCreateView)
	router.GET("articles", app.ArticleListView)
	router.GET("article_id_title", app.ArticleIDTitleListView)
	router.GET("categories", app.ArticleCategoryListView)
	router.GET("articles/detail", app.ArticleDetailByTitleView)
	router.GET("articles/calendar", app.ArticleCalendarView)
	router.GET("articles/tags", app.ArticleTagListView)
	router.PUT("articles", middleware.JwtAdmin(), app.ArticleUpdateView)
	router.DELETE("articles", middleware.JwtAdmin(), app.ArticleRemoveView)
	router.POST("articles/collections", middleware.JwtAuth(), app.ArticleCollCreateView)
	router.GET("articles/collections", middleware.JwtAuth(), app.ArticleCollListView) // 用户收藏的文章列表
	router.DELETE("articles/collections", middleware.JwtAuth(), app.ArticleCollBatchRemoveView)
	router.GET("articles/text", app.FullTextSearchView)            // 全文搜索
	router.POST("articles/digg", app.ArticleDiggView)              // 文章点赞
	router.GET("articles/content/:id", app.ArticleContentByIDView) // 文章内容
	router.GET("articles/:id", app.ArticleDetailView)              // 文章详情
	router.GET("articles/recommend", app.ArticleRecommendListView) // 推荐文章列表
}
