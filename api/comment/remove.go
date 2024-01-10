package comment

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"server/global"
	"server/models"
	"server/service/common/response"
	"server/service/redis"
	"server/utils"
	"server/utils/jwts"
)

// CommentRemoveView 删除评论
// @Tags 评论管理
// @Summary 删除评论
// @Description 删除某条评论以及子评论（如果有）
// @Param token header string  true  "token"
// @Param id path int  true  "id"
// @Router /api/comments/{id} [delete]
// @Produce json
// @Success 200 {object} response.Response{}
func (CommentApi) CommentRemoveView(c *gin.Context) {
	var cr CommentIDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	var commentModel models.CommentModel
	err = global.DB.Take(&commentModel, cr.ID).Error
	if err != nil {
		response.FailWithMessage("评论不存在", c)
		return
	}
	// 这条评论只能由当前登录人删除，或者管理员
	if !(commentModel.UserID == claims.UserID || claims.Role == 1) {
		response.FailWithMessage("权限错误，不可删除", c)
		return
	}

	// 统计评论下的子评论数 再把自己算上去
	subCommentList := models.FindAllSubCommentList(commentModel)
	count := len(subCommentList) + 1
	redis.NewCommentCount().SetCount(commentModel.ArticleID, -count)

	// 判断是否是子评论
	if commentModel.ParentCommentID != nil {
		// 子评论
		// 找父评论，减掉对应的评论数
		global.DB.Model(&models.CommentModel{}).
			Where("id = ?", *commentModel.ParentCommentID).
			Update("comment_count", gorm.Expr("comment_count - ?", count))
	}

	// 删除子评论以及当前评论
	var deleteCommentIDList []uint
	for _, model := range subCommentList {
		deleteCommentIDList = append(deleteCommentIDList, model.ID)
	}
	// 反转，然后一个一个删
	utils.Reverse(deleteCommentIDList)
	deleteCommentIDList = append(deleteCommentIDList, commentModel.ID)
	for _, id := range deleteCommentIDList {
		global.DB.Model(models.CommentModel{}).Delete("id = ?", id)
	}

	response.OkWithMessage(fmt.Sprintf("共删除 %d 条评论", len(deleteCommentIDList)), c)
	return
}
