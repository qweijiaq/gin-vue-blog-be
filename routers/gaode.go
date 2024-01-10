package routers

import (
	"server/api"
	"server/middleware"
)

func (router RouterGroup) GaodeRouter() {
	app := api.ApiGroupApp.GaodeApi
	r := router.Group("gaode")
	r.GET("weather", middleware.JwtAuth(), app.WeatherInfoView)
}
