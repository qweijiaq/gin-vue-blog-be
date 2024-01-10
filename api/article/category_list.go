package article

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"server/global"
	"server/models"
	"server/service/common/response"
)

type CategoryResponse struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

// ArticleCategoryListView 文章分类列表
// @Tags 文章管理
// @Summary 文章分类列表
// @Description 获取文章分类的列表
// @Router /api/categories [get]
// @Produce json
// @Success 200 {object} response.Response{data=[]CategoryResponse}
func (ArticleApi) ArticleCategoryListView(c *gin.Context) {
	type T struct {
		DocCountErrorUpperBound int `json:"doc_count_error_upper_bound"`
		SumOtherDocCount        int `json:"sum_other_doc_count"`
		Buckets                 []struct {
			Key      string `json:"key"`
			DocCount int    `json:"doc_count"`
		} `json:"buckets"`
	}

	agg := elastic.NewTermsAggregation().Field("category")
	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(elastic.NewBoolQuery()).
		Aggregation("category", agg).
		Size(10000).
		Do(context.Background())
	if err != nil {
		logrus.Error(err)
		return
	}
	byteData := result.Aggregations["category"]
	var categoryType T
	_ = json.Unmarshal(byteData, &categoryType)
	var categoryList = make([]CategoryResponse, 0)
	for _, i2 := range categoryType.Buckets {
		categoryList = append(categoryList, CategoryResponse{
			Label: i2.Key,
			Value: i2.Key,
		})
	}
	response.OkWithData(categoryList, c)

}
