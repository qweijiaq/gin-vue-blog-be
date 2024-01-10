package comment

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
	"server/global"
	"server/models"
	"server/service/common/response"
	"server/service/redis"
)

type CommentListRequest struct {
	ArticleID string `form:"id" uri:"id" json:"id"`
}

// CommentListView 获取文章下的评论列表
// @Tags 评论管理
// @Summary 文章评论列表
// @Description 获取文章下的评论列表
// @Param id path string  true  "id"
// @Router /api/comments/{id} [get]
// @Produce json
// @Success 200 {object} response.Response{data=[]models.CommentModel}
func (CommentApi) CommentListView(c *gin.Context) {
	var cr CommentListRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		response.FailWithError(err, &cr, c)
		return
	}
	rootCommentList := FindArticleCommentList(cr.ArticleID)

	data := filter.Select("c", rootCommentList)
	_list, _ := data.(filter.Filter)
	if string(_list.MustMarshalJSON()) == "{}" {
		list := make([]models.CommentModel, 0)
		response.OkWithList(list, 0, c)
		return
	}

	response.OkWithList(data, int64(len(rootCommentList)), c)
	return
}

func FindArticleCommentList(articleID string) (RootCommentList []*models.CommentModel) {
	// 先把文章下的根评论查出来
	global.DB.Preload("User").Find(&RootCommentList, "article_id = ? and parent_comment_id is null", articleID)
	// 遍历根评论，递归查根评论下的所有子评论
	diggInfo := redis.NewCommentDigg().GetInfo()
	for _, model := range RootCommentList {
		modelDigg := diggInfo[fmt.Sprintf("%d", model.ID)]
		model.DiggCount = model.DiggCount + modelDigg
		models.GetCommentTree(model)
	}
	return
}
