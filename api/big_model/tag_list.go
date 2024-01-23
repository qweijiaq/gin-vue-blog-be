package big_model

import (
	"github.com/gin-gonic/gin"
	"server/models"
	"server/service/common"
	"server/service/common/response"
)

type TagListResponse struct {
	models.MODEL
	Title     string `json:"title"`     // 名称
	Color     string `json:"color"`     // 颜色
	RoleCount int    `json:"roleCount"` // 角色个数
}

// TagListView 获取标签列表
func (BigModelApi) TagListView(c *gin.Context) {
	var cr models.PageInfo
	c.ShouldBindQuery(&cr)
	_list, count, _ := common.ComList(models.BigModelTagModel{}, common.Option{
		Likes:   []string{"title"},
		Preload: []string{"Roles"},
	})
	var list = make([]TagListResponse, 0)
	for _, model := range _list {
		list = append(list, TagListResponse{
			MODEL:     model.MODEL,
			Title:     model.Title,
			Color:     model.Color,
			RoleCount: len(model.Roles),
		})
	}
	response.OkWithList(list, count, c)
}
