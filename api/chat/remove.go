package chat

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/service/common/response"
)

// ChatRemoveView 删除群聊的聊天记录
// @Tags 聊天管理
// @Summary 删除群聊的聊天记录
// @Description 删除群聊的聊天记录
// @Param data body models.RemoveRequest   false  "表示多个参数"
// @Router /api/chat_groups_records [delete]
// @Produce json
// @Success 200 {object} response.Response{}
func (ChatApi) ChatRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}

	var chatList []models.ChatModel
	global.DB.Find(&chatList, cr.IDList)

	if len(chatList) > 0 {
		err = global.DB.Delete(&chatList).Error
		if err != nil {
			response.FailWithMessage("群聊记录删除失败", c)
			return
		}
	}

	response.OkWithMessage(fmt.Sprintf("共删除记录%d条", len(chatList)), c)

}
