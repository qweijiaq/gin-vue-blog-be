package routers

import "server/api"

func (router RouterGroup) NewsRouter() {
	app := api.ApiGroupApp.NewsApi
	router.POST("news", app.NewListView)
}
