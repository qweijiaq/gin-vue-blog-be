package message

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/service/common/response"
	"server/utils/jwts"
)

type MessageRequest struct {
	RevUserID uint   `json:"rev_user_id" binding:"required"` // 接收人id
	Content   string `json:"content" binding:"required"`     // 消息内容
}

// MessageCreateView 发布消息
// @Tags 消息管理
// @Summary 发布消息
// @Description 发布消息
// @Param data body MessageRequest  true  "查询参数"
// @Param token header string  true  "token"
// @Router /api/messages [post]
// @Produce json
// @Success 200 {object} response.Response{}
func (MessageApi) MessageCreateView(c *gin.Context) {
	// 当前用户发布消息
	// SendUserID 就是当前登录人的id
	var cr MessageRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithError(err, &cr, c)
		return
	}
	var senUser, revUser models.UserModel

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	err = global.DB.Take(&senUser, claims.UserID).Error
	if err != nil {
		response.FailWithMessage("发送人不存在", c)
		return
	}
	err = global.DB.Take(&revUser, cr.RevUserID).Error
	if err != nil {
		response.FailWithMessage("接收人不存在", c)
		return
	}

	err = global.DB.Create(&models.MessageModel{
		SendUserID:       senUser.ID,
		SendUserNickName: senUser.NickName,
		SendUserAvatar:   senUser.Avatar,
		RevUserID:        cr.RevUserID,
		RevUserNickName:  revUser.NickName,
		RevUserAvatar:    revUser.Avatar,
		IsRead:           false,
		Content:          cr.Content,
	}).Error
	if err != nil {
		global.Log.Error(err)
		response.FailWithMessage("消息发送失败", c)
		return
	}
	response.OkWithMessage("消息发送成功", c)
	return
}
