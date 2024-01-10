package routers

import (
	"server/api"
)

func (router RouterGroup) FeedbackRouter() {
	app := api.ApiGroupApp.FeedbackApi
	router.POST("feedbacks", app.FeedBackCreateView)
	router.GET("feedbacks", app.FeedBackListView)
	router.DELETE("feedbacks", app.FeedBackRemoveView)
}
