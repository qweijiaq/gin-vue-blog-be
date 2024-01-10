package article

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/service/common/response"
	"server/service/es"
	"server/service/redis"
	"server/utils/jwts"
)

type ArticleItem struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}
type ArticleDetailResponse struct {
	models.ArticleModel
	IsCollect bool         `json:"is_collect"` // 用户是否收藏文章
	Next      *ArticleItem `json:"next"`       // 上一篇文章
	Prev      *ArticleItem `json:"prev"`       // 下一篇文章
}

// ArticleDetailView 文章详情
// @Tags 文章管理
// @Summary 文章详情
// @Description 查看文章详情内容
// @Param id path string  true  "id"
// @Router /api/articles/{id} [get]
// @Produce json
// @Success 200 {object} response.Response{data=models.ArticleModel}
func (ArticleApi) ArticleDetailView(c *gin.Context) {
	var cr models.ESIDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}
	redis.NewArticleLook().Set(cr.ID)
	model, err := es.CommDetail(cr.ID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	isCollect := IsUserArticleColl(c, model.ID)
	var articleDetail = ArticleDetailResponse{
		ArticleModel: model,
		IsCollect:    isCollect,
	}
	// 查上一篇  下一篇文章
	// 根据分类，查文章列表，然后找当前文章所在的位置
	list, _, err := es.CommList(es.Option{
		PageInfo: models.PageInfo{
			Limit: 10000,
			Page:  1,
		},
		Category: model.Category,
	})
	if err == nil {
		var currentIndex = -1

		// 查找当前记录的索引
		for i, item := range list {
			if item.ID == model.ID {
				currentIndex = i
				break
			}
		}

		var previousModel models.ArticleModel
		var nextModel models.ArticleModel

		if currentIndex > 0 {
			previousModel = list[currentIndex-1]
			articleDetail.Next = &ArticleItem{
				ID:    previousModel.ID,
				Title: previousModel.Title,
			}
		}

		// 查找下一个记录
		if currentIndex < len(list)-1 {
			nextModel = list[currentIndex+1]
			articleDetail.Prev = &ArticleItem{
				ID:    nextModel.ID,
				Title: nextModel.Title,
			}
		}
	}

	response.OkWithData(articleDetail, c)
}

func IsUserArticleColl(c *gin.Context, articleID string) (isCollect bool) {
	// 判断用户是否正常登录
	token := c.GetHeader("token")
	if token == "" {
		return
	}
	claims, err := jwts.ParseToken(token)
	if err != nil {
		return
	}
	// 判断是否在redis中
	if redis.CheckLogout(token) {
		return
	}
	var count int64
	global.DB.Model(models.UserCollectModel{}).Where("user_id = ? and article_id =?", claims.UserID, articleID).Count(&count)
	if count == 0 {
		return
	}
	return true
}

type ArticleDetailRequest struct {
	Title string `json:"title" form:"title"`
}

func (ArticleApi) ArticleDetailByTitleView(c *gin.Context) {
	var cr ArticleDetailRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}
	model, err := es.CommDetailByKeyword(cr.Title)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(model, c)
}
