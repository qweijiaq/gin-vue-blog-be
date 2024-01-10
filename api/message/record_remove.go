package message

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/service/common/response"
)

// MessageRecordRemoveView 删除用户的消息记录
// @Tags 消息管理
// @Summary 删除用户的消息记录
// @Description 删除用户的消息记录
// @Router /api/message_users [delete]
// @Param token header string  true  "token"
// @Param data body models.RemoveRequest   true  "查询参数"
// @Produce json
// @Success 200 {object} response.Response{]}
func (MessageApi) MessageRecordRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}

	var messageList []models.MessageModel
	global.DB.Find(&messageList, cr.IDList)

	if len(messageList) > 0 {
		err = global.DB.Delete(&messageList).Error
		if err != nil {
			response.FailWithMessage("消息记录删除失败", c)
			return
		}
	}

	response.OkWithMessage(fmt.Sprintf("共删除记录%d条", len(messageList)), c)
}
