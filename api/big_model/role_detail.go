package big_model

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/service/common/response"
)

type RoleDetailResponse struct {
	models.MODEL
	Icon      string        `json:"icon"`
	Name      string        `json:"name"`
	Abstract  string        `json:"abstract"`
	Tags      []TagResponse `json:"tags"`
	ChatCount int64         `json:"chatCount"`
}

type TagResponse struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Color string `json:"color"`
}

// RoleDetailView 角色详情
func (BigModelApi) RoleDetailView(c *gin.Context) {
	var cr models.IDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		response.FailWithValidError(err, c)
		return
	}
	sessionId := c.Query("sessionId")
	var arm models.BigModelRoleModel
	err = global.DB.Preload("Tags").Take(&arm, cr.ID).Error
	if err != nil {
		response.FailWithMessage("角色不存在", c)
		return
	}

	var tags = make([]TagResponse, 0)
	for _, tag := range arm.Tags {
		tags = append(tags, TagResponse{
			ID:    tag.ID,
			Title: tag.Title,
			Color: tag.Color,
		})
	}
	res := RoleDetailResponse{
		MODEL:    arm.MODEL,
		Icon:     arm.Icon,
		Name:     arm.Name,
		Abstract: arm.Abstract,
		Tags:     tags,
	}

	// 找这个角色进行了多少次对话
	global.DB.Model(models.BigModelChatModel{}).Where("role_id = ? and session_id = ?", cr.ID, sessionId).Count(&res.ChatCount)

	response.OkWithData(res, c)
}
