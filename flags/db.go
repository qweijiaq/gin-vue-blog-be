package flags

import (
	"server/global"
	"server/models"
	"server/plugins/log_stash"
	logStashV2 "server/plugins/log_stash_v2"
)

func DB() {
	var err error
	global.DB.SetupJoinTable(&models.MenuModel{}, "Banners", &models.MenuBannerModel{})
	// 生成四张表的表结构
	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&models.BannerModel{},
			&models.TagModel{},
			&models.MessageModel{},
			&models.AdvertModel{},
			&models.UserModel{},
			&models.CommentModel{},
			&models.UserCollectModel{},
			&models.MenuModel{},
			&models.MenuBannerModel{},
			&models.ReplyModel{},
			&models.LoginDataModel{},
			&models.ChatModel{},
			&models.FeedbackModel{},
			&log_stash.LogStashModel{},
			&logStashV2.LogModel{},
			&models.UserScopeModel{},
			&models.AutoReplyModel{},
			&models.BigModelRoleModel{},
			&models.BigModelTagModel{},
			&models.BigModelChatModel{},
			&models.BigModelSessionModel{},
		)
	if err != nil {
		global.Log.Error("[ error ] 生成数据库表结构失败！")
		return
	}
	global.Log.Info("[ success ] 生成数据库表结构成功！")
}
