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

// FullTextSearchView 全文搜索列表
// @Tags 文章管理
// @Summary 全文搜索列表
// @Description 获取全文搜索时返回的列表
// @Param data query models.PageInfo   false  "表示多个参数"
// @Router /api/articles/text [get]
// @Produce json
// @Success 200 {object} response.Response{data=response.ListResponse[models.FullTextModel]}
func (ArticleApi) FullTextSearchView(c *gin.Context) {
	var cr models.PageInfo
	_ = c.ShouldBindQuery(&cr)

	boolQuery := elastic.NewBoolQuery()

	if cr.Key != "" {
		boolQuery.Must(elastic.NewMultiMatchQuery(cr.Key, "title", "body"))
	}
	if cr.Page == 0 {
		cr.Page = 1
	}
	if cr.Limit == 0 {
		cr.Limit = 10
	}
	from := (cr.Page - 1) * cr.Limit

	result, err := global.ESClient.
		Search(models.FullTextModel{}.Index()).
		Query(boolQuery).
		Highlight(elastic.NewHighlight().Field("body")).
		From(from).
		Size(cr.Limit).
		Do(context.Background())
	if err != nil {
		return
	}
	count := result.Hits.TotalHits.Value // 搜索到结果总条数
	fullTextList := make([]models.FullTextModel, 0)
	for _, hit := range result.Hits.Hits {
		var model models.FullTextModel
		json.Unmarshal(hit.Source, &model)

		body, ok := hit.Highlight["body"]
		if ok {
			model.Body = body[0]
		}

		fullTextList = append(fullTextList, model)
	}

	response.OkWithList(fullTextList, count, c)
}
