package routers

import (
	"server/api"
	"server/middleware"
)

func (router RouterGroup) TagRouter() {
	app := api.ApiGroupApp.TagApi
	router.POST("tags", middleware.JwtAdmin(), app.TagCreateView)
	router.GET("tags", app.TagListView)
	router.GET("tag_names", app.TagNameListView)
	router.PUT("tags/:id", middleware.JwtAdmin(), app.TagUpdateView)
	router.DELETE("tags", middleware.JwtAdmin(), app.TagRemoveView)
}
