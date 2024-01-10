package article

import (
	"github.com/gin-gonic/gin"
	"server/models"
	"server/service/common/response"
	"server/service/redis"
)

// ArticleDiggView 文章点赞
// @Tags 文章管理
// @Summary 文章点赞
// @Description 为某一篇文章点赞
// @Param data body models.ESIDRequest   true  "表示多个参数"
// @Router /api/articles/digg [post]
// @Produce json
// @Success 200 {object} response.Response{}
func (ArticleApi) ArticleDiggView(c *gin.Context) {
	var cr models.ESIDRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}
	// 对长度校验
	// 查 ES 数据库
	redis.NewDigg().Set(cr.ID)
	response.OkWithMessage("文章点赞成功", c)
}
