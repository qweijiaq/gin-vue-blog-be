package api

import (
	"server/api/advert"
	"server/api/article"
	"server/api/big_model"
	"server/api/chat"
	"server/api/comment"
	"server/api/data"
	"server/api/feedback"
	"server/api/gaode"
	"server/api/images"
	"server/api/log"
	"server/api/log_v2"
	"server/api/menu"
	"server/api/message"
	"server/api/news"
	"server/api/role"
	"server/api/settings"
	"server/api/tag"
	"server/api/user"
)

type ApiGroup struct {
	SettingsApi settings.SettingsApi
	ImagesApi   images.ImagesApi
	AdvertApi   advert.AdvertApi
	MenuApi     menu.MenuApi
	UserApi     user.UserApi
	TagApi      tag.TagApi
	MessageApi  message.MessageApi
	ArticleApi  article.ArticleApi
	CommentApi  comment.CommentApi
	NewsApi     news.NewsApi
	ChatApi     chat.ChatApi
	LogApi      log.LogApi
	DataApi     data.DataApi
	LogV2Api    log_v2.LogApi
	RoleApi     role.RoleApi
	GaodeApi    gaode.GaodeApi
	FeedbackApi feedback.FeedbackApi
	BigModelApi big_model.BigModelApi
}

var ApiGroupApp = new(ApiGroup)
