package cron

import (
	"gorm.io/gorm"
	"server/global"
	"server/models"
	"server/service/redis"
)

// SyncCommentData 同步评论数据到数据库
func SyncCommentData() {
	commentDiggInfo := redis.NewCommentDigg().GetInfo()
	for key, count := range commentDiggInfo {
		var comment models.CommentModel
		err := global.DB.Take(&comment, key).Error
		if err != nil {
			global.Log.Error(err)
			continue
		}
		err = global.DB.Model(&comment).
			Update("digg_count", gorm.Expr("digg_count + ?", count)).Error
		if err != nil {
			global.Log.Error(err)
			continue
		}
		global.Log.Infof("%s 更新成功 新的点赞数为：%d", comment.Content, comment.DiggCount)
	}
	redis.NewCommentDigg().Clear()
}
