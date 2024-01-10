package article

import (
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
	"github.com/olivere/elastic/v7"
	"server/global"
	"server/models"
	"server/service/common/response"
	"server/service/es"
	"server/service/redis"
	"server/utils/jwts"
	"time"
)

type ArticleSearchRequest struct {
	models.PageInfo
	Tag      string `json:"tag" form:"tag"`
	Category string `json:"category" form:"category"`
	IsUser   bool   `json:"is_user" form:"is_user"` // 根据这个参数判断是否显示我收藏的文章列表
	Date     string `json:"date" form:"date"`       // 发布时间搜索
}

// ArticleListView 文章列表
// @Tags 文章管理
// @Summary 文章列表
// @Description 获取文章列表
// @Param data query ArticleSearchRequest   false  "表示多个参数"
// @Param token header string  false  "token"
// @Router /api/articles [get]
// @Produce json
// @Success 200 {object} response.Response{data=response.ListResponse[models.ArticleModel]}
func (ArticleApi) ArticleListView(c *gin.Context) {
	var cr ArticleSearchRequest
	if err := c.ShouldBindQuery(&cr); err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}
	boolSearch := elastic.NewBoolQuery()

	if cr.IsUser {
		token := c.GetHeader("token")
		claims, err := jwts.ParseToken(token)
		if err == nil && !redis.CheckLogout(token) {
			boolSearch.Must(elastic.NewTermsQuery("user_id", claims.UserID))
		}
	}

	if cr.Date != "" {
		date, err := time.Parse("2006-01-02", cr.Date)
		if err == nil {
			boolSearch.Must(elastic.NewRangeQuery("created_at").
				Gte(date.Format("2006-01-02") + " 00:00:00").
				Lte(date.Format("2006-01-02") + " 23:59:59"))
		}
	}

	list, count, err := es.CommList(es.Option{
		PageInfo: cr.PageInfo,
		Fields:   []string{"title", "content"},
		Tag:      cr.Tag,
		Query:    boolSearch,
		Category: cr.Category,
	})
	if err != nil {
		global.Log.Error(err)
		response.OkWithMessage("查询失败", c)
		return
	}

	// json-filter空值问题
	data := filter.Omit("list", list)
	_list, _ := data.(filter.Filter)
	if string(_list.MustMarshalJSON()) == "{}" {
		list = make([]models.ArticleModel, 0)
		response.OkWithList(list, int64(count), c)
		return
	}
	response.OkWithList(data, int64(count), c)
}
