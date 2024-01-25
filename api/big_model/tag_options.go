package big_model

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/service/common/response"
)

// TagOptionsListView 获取标签 ID 列表
func (BigModelApi) TagOptionsListView(c *gin.Context) {
	var list []models.Options
	global.DB.Model(models.BigModelTagModel{}).Select("id as value", "title as label").Scan(&list)

	response.OkWithData(list, c)
}
