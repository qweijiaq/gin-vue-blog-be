package article

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/service/common/response"
	"server/service/es"
	"server/utils/jwts"
)

// ArticleCollCreateView 用户收藏或取消收藏文章
// @Tags 文章管理
// @Summary 收藏或取消收藏
// @Description 用户收藏或取消收藏文章
// @Param data body models.ESIDRequest   true  "表示多个参数"
// @Param token header string  true  "token"
// @Router /api/articles/collect [post]
// @Produce json
// @Success 200 {object} response.Response{}
func (ArticleApi) ArticleCollCreateView(c *gin.Context) {
	var cr models.ESIDRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	model, err := es.CommDetail(cr.ID)
	if err != nil {
		response.FailWithMessage("文章不存在", c)
		return
	}

	var coll models.UserCollectModel
	err = global.DB.Take(&coll, "user_id = ? and article_id = ?", claims.UserID, cr.ID).Error
	var num = -1
	if err != nil {
		// 没有找到 收藏文章
		global.DB.Create(&models.UserCollectModel{
			UserID:    claims.UserID,
			ArticleID: cr.ID,
		})
		// 给文章的收藏数 +1
		num = 1
	}
	// 取消收藏
	// 文章数 -1
	global.DB.Delete(&coll)

	// 更新文章收藏数
	err = es.ArticleUpdate(cr.ID, map[string]any{
		"collects_count": model.CollectsCount + num,
	})
	if num == 1 {
		response.OkWithMessage("收藏文章成功", c)
	} else {
		response.OkWithMessage("取消收藏成功", c)
	}
}
