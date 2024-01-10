package cron

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"server/global"
	"server/models"
	"server/service/redis"
)

// SyncArticleData 同步文章数据到es
func SyncArticleData() {
	// 1.查询es中的全部数据，为后面的数据更新做准备
	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(elastic.NewMatchAllQuery()).
		Size(10000).
		Do(context.Background())
	if err != nil {
		global.Log.Error(err)
		return
	}
	// 2.拿到redis中的缓存数据
	diggInfo := redis.NewDigg().GetInfo()
	lookInfo := redis.NewArticleLook().GetInfo()
	commentInfo := redis.NewCommentCount().GetInfo()

	for _, hit := range result.Hits.Hits {
		var article models.ArticleModel
		err = json.Unmarshal(hit.Source, &article)
		if err != nil {
			global.Log.Error(err)
			continue
		}
		digg := diggInfo[hit.Id]
		look := lookInfo[hit.Id]
		comment := commentInfo[hit.Id]
		// 3. 计算新的数据   之前的数据+缓存中的数据
		newDigg := article.DiggCount + digg
		newLook := article.LookCount + look
		newComment := article.CommentCount + comment
		// 4. 要判断一下，是否有变化，如果三个变化之后的数据和之前的是一样的
		if digg == 0 && look == 0 && comment == 0 {
			global.Log.Infof("%s无变化", article.Title)
			continue
		}
		// 5. 更新
		_, err = global.ESClient.Update().
			Index(models.ArticleModel{}.Index()).
			Id(hit.Id).
			Doc(map[string]int{
				"look_count":    newLook,
				"comment_count": newComment,
				"digg_count":    newDigg,
			}).Do(context.Background())
		if err != nil {
			global.Log.Error(err)
			continue
		}
		global.Log.Infof("%s 更新成功 点赞数：%d 评论数：%d  浏览量：%d",
			article.Title, newDigg, newComment, newLook)
	}
	// 6. 清除redis中的数据
	redis.NewDigg().Clear()
	redis.NewArticleLook().Clear()
	redis.NewCommentCount().Clear()
}
