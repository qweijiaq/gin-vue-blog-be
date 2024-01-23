package big_model

import (
	"github.com/gin-gonic/gin"
	"server/models"
	"server/service/common"
	"server/service/common/response"
)

// RoleListView 获取角色列表
func (BigModelApi) RoleListView(c *gin.Context) {
	var cr models.PageInfo
	c.ShouldBindQuery(&cr)

	list, count, _ := common.ComList(models.BigModelRoleModel{}, common.Option{
		PageInfo: cr,
		Likes:    []string{"name"},
	})

	response.OkWithList(list, count, c)
	return
}
