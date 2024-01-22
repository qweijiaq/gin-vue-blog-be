package routers

import (
	"server/api"
	"server/middleware"
)

func (router RouterGroup) BigModelRouter() {
	app := api.ApiGroupApp.BigModelApi
	router.GET("bigModel/usable", middleware.JwtAdmin(), app.ModelUsableListView)
	router.GET("bigModel/setting", app.ModelSettingView)
}
