package big_model

import (
	"github.com/gin-gonic/gin"
	"server/models"
	"server/service/common"
	"server/service/common/response"
	"server/utils/jwts"
	"time"
)

type RoleSessionsRequest struct {
	models.PageInfo
	RoleID uint `json:"roleID" form:"roleID" binding:"required"`
}

type RoleSessionResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
}

// RoleSessionListView 角色会话列表
func (BigModelApi) RoleSessionListView(c *gin.Context) {
	var cr RoleSessionsRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		response.FailWithValidError(err, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	_list, count, _ := common.ComList(models.BigModelSessionModel{UserID: claims.UserID, RoleID: cr.RoleID}, common.Option{
		PageInfo: cr.PageInfo,
		Likes:    []string{"name"},
	})
	var list = make([]RoleSessionResponse, 0)
	for _, model := range _list {
		list = append(list, RoleSessionResponse{
			ID:        model.ID,
			CreatedAt: model.CreatedAt,
			Name:      model.Name,
		})
	}
	response.OkWithList(list, count, c)
}
