package big_model

import (
	"github.com/gin-gonic/gin"
	"server/models"
	"server/service/common"
	"server/service/common/response"
)

// AutoReplyListView 获取自动回复规则列表
func (BigModelApi) AutoReplyListView(c *gin.Context) {
	var cr models.PageInfo
	c.ShouldBindQuery(&cr)

	list, count, _ := common.ComList(models.AutoReplyModel{}, common.Option{
		PageInfo: cr,
		Likes:    []string{"name"},
	})

	response.OkWithList(list, count, c)
	return
}
