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

type TagsResponse struct {
	Tag           string   `json:"tag"`
	Count         int      `json:"count"`
	ArticleIDList []string `json:"article_id_list"`
	CreatedAt     string   `json:"created_at"`
}

type TagsType struct {
	DocCountErrorUpperBound int `json:"doc_count_error_upper_bound"`
	SumOtherDocCount        int `json:"sum_other_doc_count"`
	Buckets                 []struct {
		Key      string `json:"key"`
		DocCount int    `json:"doc_count"`
		Articles struct {
			DocCountErrorUpperBound int `json:"doc_count_error_upper_bound"`
			SumOtherDocCount        int `json:"sum_other_doc_count"`
			Buckets                 []struct {
				Key      string `json:"key"`
				DocCount int    `json:"doc_count"`
			} `json:"buckets"`
		} `json:"articles"`
	} `json:"buckets"`
}

// ArticleTagListView 文章标签列表
// @Tags 文章管理
// @Summary 文章标签列表
// @Description 获取文章对应的标签列表
// @Param data query models.PageInfo   false  "表示多个参数"
// @Router /api/articles/tags [get]
// @Produce json
// @Success 200 {object} response.Response{data=response.ListResponse[TagsResponse]}
func (ArticleApi) ArticleTagListView(c *gin.Context) {

	var cr models.PageInfo
	_ = c.ShouldBindQuery(&cr)

	if cr.Limit == 0 {
		cr.Limit = 10
	}
	offset := (cr.Page - 1) * cr.Limit
	if offset < 0 {
		offset = 0
	}

	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Aggregation("tags", elastic.NewCardinalityAggregation().Field("tags")).
		Size(0).
		Do(context.Background())
	cTag, _ := result.Aggregations.Cardinality("tags")
	count := int64(*cTag.Value)

	agg := elastic.NewTermsAggregation().Field("tags")

	agg.SubAggregation("articles", elastic.NewTermsAggregation().Field("keyword"))
	agg.SubAggregation("page", elastic.NewBucketSortAggregation().From(offset).Size(cr.Limit))

	query := elastic.NewBoolQuery()

	result, err = global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(query).
		Aggregation("tags", agg).
		Size(0).
		Do(context.Background())
	if err != nil {
		global.Log.Error(err)
		response.FailWithMessage(err.Error(), c)
		return
	}
	var tagType TagsType
	var tagList = make([]*TagsResponse, 0)
	_ = json.Unmarshal(result.Aggregations["tags"], &tagType)
	var tagStringList []string
	for _, bucket := range tagType.Buckets {

		var articleList []string
		for _, s := range bucket.Articles.Buckets {
			articleList = append(articleList, s.Key)
		}

		tagList = append(tagList, &TagsResponse{
			Tag:           bucket.Key,
			Count:         bucket.DocCount,
			ArticleIDList: articleList,
		})
		tagStringList = append(tagStringList, bucket.Key)
	}

	var tagModelList []models.TagModel
	global.DB.Find(&tagModelList, "title in ?", tagStringList)
	var tagDate = map[string]string{}
	for _, model := range tagModelList {
		tagDate[model.Title] = model.CreatedAt.Format("2006-01-02 15:04:05")
	}
	for _, response := range tagList {
		response.CreatedAt = tagDate[response.Tag]
	}
	response.OkWithList(tagList, count, c)
}
