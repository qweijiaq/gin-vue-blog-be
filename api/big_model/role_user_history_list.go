package big_model

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/service/common/response"
	"server/utils/jwts"
)

// RoleUserHistoryListView 用户角色历史列表
func (BigModelApi) RoleUserHistoryListView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	var roleIdList []uint
	global.DB.Model(models.BigModelSessionModel{}).Where("user_id = ?", claims.UserID).Group("role_id").Select("role_id").Scan(&roleIdList)
	var roleList []models.BigModelRoleModel
	global.DB.Order("created_at desc").Find(&roleList, "id in ?", roleIdList)

	var list = make([]RoleItem, 0)
	for _, model := range roleList {
		list = append(list, RoleItem{
			ID:       model.ID,
			Name:     model.Name,
			Abstract: model.Abstract,
			Icon:     model.Icon,
		})
	}
	response.OkWithData(list, c)
	return
}
