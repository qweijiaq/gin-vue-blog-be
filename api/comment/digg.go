package comment

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/service/common/response"
	"server/service/redis"
)

type CommentIDRequest struct {
	ID uint `json:"id" form:"id" uri:"id"`
}

// CommentDigg 点赞某一个评论
// @Tags 评论管理
// @Summary 评论点赞
// @Description 点赞某一个评论
// @Param id path int  true  "id"
// @Router /api/comments/digg/{id} [get]
// @Produce json
// @Success 200 {object} response.Response{}
func (CommentApi) CommentDigg(c *gin.Context) {
	var cr CommentIDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}

	var commentModel models.CommentModel
	err = global.DB.Take(&commentModel, cr.ID).Error
	if err != nil {
		response.FailWithMessage("评论不存在", c)
		return
	}

	redis.NewCommentDigg().Set(fmt.Sprintf("%d", cr.ID))

	response.OkWithMessage("评论点赞成功", c)
	return

}
