package article

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/ctype"
	"server/service/common/response"
	"server/service/es"
	"time"
)

type ArticleUpdateRequest struct {
	Title       string   `json:"title"`        // 文章标题
	Abstract    string   `json:"abstract"`     // 文章简介
	Content     string   `json:"content"`      // 文章内容
	Category    string   `json:"category"`     // 文章分类
	Source      string   `json:"source"`       // 文章来源
	Link        string   `json:"link"`         // 原文链接
	BannerID    uint     `json:"banner_id"`    // 文章封面id
	Tags        []string `json:"tags"`         // 文章标签
	IsRecommend bool     `json:"is_recommend"` // 是否推荐
	ID          string   `json:"id"`
}

// ArticleUpdateView 文章更新
// @Tags 文章管理
// @Summary 文章更新
// @Description 更新某篇文章的内容
// @Param data body ArticleUpdateRequest   false  "传什么参数更新什么，不传不更"
// @Param token header string  true  "token"
// @Router /api/articles [put]
// @Produce json
// @Success 200 {object} response.Response{}
func (ArticleApi) ArticleUpdateView(c *gin.Context) {
	var cr ArticleUpdateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Error(err)
		response.FailWithError(err, &cr, c)
		return
	}
	var bannerUrl string
	if cr.BannerID != 0 {
		err = global.DB.Model(models.BannerModel{}).Where("id = ?", cr.BannerID).Select("path").Scan(&bannerUrl).Error
		if err != nil {
			response.FailWithMessage("banner不存在", c)
			return
		}
	}

	article := models.ArticleModel{
		UpdatedAt:   time.Now().Format("2006-01-02 15:04:05"),
		Title:       cr.Title,
		Keyword:     cr.Title,
		Abstract:    cr.Abstract,
		Content:     cr.Content,
		Category:    cr.Category,
		Source:      cr.Source,
		Link:        cr.Link,
		BannerID:    cr.BannerID,
		BannerUrl:   bannerUrl,
		Tags:        cr.Tags,
		IsRecommend: cr.IsRecommend,
	}

	maps := structs.Map(&article)
	var DataMap = map[string]any{}
	// 去掉空值
	for key, v := range maps {
		switch val := v.(type) {
		case string:
			if val == "" {
				continue
			}
		case uint:
			if val == 0 {
				continue
			}
		case int:
			if val == 0 {
				continue
			}
		case ctype.Array:
			if len(val) == 0 {
				continue
			}
		case []string:
			if len(val) == 0 {
				continue
			}
		case bool:
			if val == false {
				continue
			}
		}

		DataMap[key] = v
	}

	err = article.GetDataByID(cr.ID)
	if err != nil {
		global.Log.Error(err)
		response.FailWithMessage("文章不存在", c)
		return
	}

	err = es.ArticleUpdate(cr.ID, DataMap)
	if err != nil {
		global.Log.Error(err)
		response.FailWithMessage("文章更新失败", c)
		return
	}

	// 更新成功，同步数据到全文搜索
	newArticle, _ := es.CommDetail(cr.ID)
	if article.Content != newArticle.Content || article.Title != newArticle.Title {
		es.DeleteFullTextByArticleID(cr.ID)
		es.AsyncArticleByFullText(cr.ID, article.Title, newArticle.Content)
	}

	response.OkWithMessage("更新成功", c)
}
