package big_model

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"server/global"
	"server/models"
	"server/service/common/response"
	"server/utils/jwts"
)

// ChatUserRemoveView 用户删除对话
func (BigModelApi) ChatUserRemoveView(c *gin.Context) {
	var cr models.IDRequest
	// control request
	err := c.ShouldBindUri(&cr)
	if err != nil {
		response.FailWithValidError(err, c)
		return
	}

	// 找会话
	var chat models.BigModelChatModel
	err = global.DB.Take(&chat, cr.ID).Error
	if err != nil {
		response.FailWithMessage("对话不存在", c)
		return
	}

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	if chat.UserID != claims.UserID {
		response.FailWithMessage("对话鉴权失败", c)
		return
	}
	// 删除会话
	err = global.DB.Delete(&chat).Error
	if err != nil {
		logrus.Error(err)
		response.FailWithMessage("对话删除失败", c)
		return
	}
	response.OkWithMessage("对话删除成功", c)
}
