package message

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/service/common"
	"server/service/common/response"
	"server/utils/jwts"
)

type MessageUserRecordByMeRequest struct {
	models.PageInfo
	UserID uint `json:"userID" form:"userID" binding:"required"`
}

// MessageUserRecordByMeView 我与某个用户的聊天列表
// @Tags 消息管理
// @Summary 我与某个用户的聊天列表
// @Description 我与某个用户的聊天列表
// @Router /api/message_users/record/me [get]
// @Param token header string  true  "token"
// @Param data query MessageUserRecordByMeRequest  true  "参数"
// @Produce json
// @Success 200 {object} response.Response{data=response.ListResponse[models.MessageModel]}
func (m MessageApi) MessageUserRecordByMeView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	var cr MessageUserRecordByMeRequest
	c.ShouldBindQuery(&cr)

	cr.Sort = "created_at asc"
	list, count, _ := common.ComList(models.MessageModel{}, common.Option{
		PageInfo: cr.PageInfo,
		Where:    global.DB.Where("(send_user_id = ? and rev_user_id = ? ) or ( rev_user_id = ? and send_user_id = ? )", claims.UserID, cr.UserID, claims.UserID, cr.UserID),
	})

	response.OkWithList(list, count, c)
}
