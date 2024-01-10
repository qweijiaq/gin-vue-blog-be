package article

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"server/global"
	"server/models"
	"server/service/common/response"
)

type ArticleRecommendSearchRequest struct {
	models.PageInfo
}

type recommendArticle struct {
	ID           string `json:"id"`
	BannerUrl    string `json:"banner_url"`
	Title        string `json:"title"`
	Content      string `json:"content"`
	Abstract     string `json:"abstract"`
	CreatedAt    string `json:"created_at"`
	LookCount    int    `json:"look_count"`
	CommentCount int    `json:"comment_count"`
	DiggCount    int    `json:"digg_count"`
}

// ArticleRecommendListView 推荐文章列表
// @Tags 文章管理
// @Summary 推荐文章列表
// @Description 获取推荐文章列表
// @Param data query ArticleRecommendSearchRequest   false  "表示多个参数"
// @Param token header string  false  "token"
// @Router /api/articles [get]
// @Produce json
// @Success 200 {object} response.Response{data=response.ListResponse[models.ArticleModel]}
func (ArticleApi) ArticleRecommendListView(c *gin.Context) {
	// 构造Elasticsearch查询
	query := elastic.NewBoolQuery().Must(elastic.NewTermQuery("is_recommend", true))

	// 执行Elasticsearch查询
	searchResult, err := global.ESClient.Search().
		Index("article_index").
		Query(query).
		Size(10).
		Do(context.Background())

	if err != nil {
		global.Log.Error(err)
		response.OkWithMessage("查询失败", c)
		return
	}

	var recommendedArticles []recommendArticle

	// 处理 Elasticsearch 查询结果
	for _, hit := range searchResult.Hits.Hits {
		var article recommendArticle
		if err := json.Unmarshal(hit.Source, &article); err != nil {
			global.Log.Error(err)
			response.OkWithMessage("处理查询结果失败", c)
			return
		}
		article.ID = hit.Id
		recommendedArticles = append(recommendedArticles, article)
	}

	response.OkWithList(recommendedArticles, int64(len(recommendedArticles)), c)
}
