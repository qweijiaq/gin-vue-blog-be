package comment

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"server/global"
	"server/models"
	"server/service/common/response"
	"server/service/es"
	"server/service/redis"
	"server/utils/jwts"
)

type CommentRequest struct {
	ArticleID       string `json:"article_id" binding:"required" msg:"请选择文章"`
	Content         string `json:"content" binding:"required" msg:"请输入评论内容"`
	ParentCommentID *uint  `json:"parent_comment_id"` // 父评论id
}

// CommentCreateView 发布评论
// @Tags 评论管理
// @Summary 发布评论
// @Description 对某篇文章发布评论
// @Param data body CommentRequest   true  "表示多个参数"
// @Param token header string  true  "token"
// @Router /api/comments [post]
// @Produce json
// @Success 200 {object} response.Response{}
func (CommentApi) CommentCreateView(c *gin.Context) {
	var cr CommentRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithError(err, &cr, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	// 文章是否存在
	_, err = es.CommDetail(cr.ArticleID)
	if err != nil {
		response.FailWithMessage("文章不存在", c)
		return
	}

	// 判断是否是子评论
	if cr.ParentCommentID != nil {
		// 子评论
		// 给父评论数 +1
		// 父评论id
		var parentComment models.CommentModel
		// 找父评论
		err = global.DB.Take(&parentComment, cr.ParentCommentID).Error
		if err != nil {
			response.FailWithMessage("父评论不存在", c)
			return
		}
		// 判断父评论的文章是否和当前文章一致
		if parentComment.ArticleID != cr.ArticleID {
			response.FailWithMessage("评论文章不一致", c)
			return
		}
		// 给父评论评论数+1
		global.DB.Model(&parentComment).Update("comment_count", gorm.Expr("comment_count + 1"))
	}
	// 添加评论
	global.DB.Create(&models.CommentModel{
		ParentCommentID: cr.ParentCommentID,
		Content:         cr.Content,
		ArticleID:       cr.ArticleID,
		UserID:          claims.UserID,
	})
	// 拿到文章数，新的文章评论数存缓存里
	//newCommentCount := article.CommentCount + 1
	// 给文章评论数 +1
	redis.NewCommentCount().Set(cr.ArticleID)
	response.OkWithMessage("文章评论成功", c)
	return
}
