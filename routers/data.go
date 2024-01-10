package routers

import "server/api"

func (router RouterGroup) DataRouter() {
	app := api.ApiGroupApp.DataApi
	router.GET("data_login", app.SevenLogin)
	router.GET("data_sum", app.DataSumView)
}
