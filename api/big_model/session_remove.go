package big_model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"server/global"
	"server/models"
	"server/service/common/response"
)

// SessionRemoveView 管理员删除会话
func (BigModelApi) SessionRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithValidError(err, c)
		return
	}

	var list []models.BigModelSessionModel
	count := global.DB.Preload("ChatList").Find(&list, cr.IDList).RowsAffected
	if count == 0 {
		response.FailWithMessage("记录不存在", c)
		return
	}

	if len(list) > 0 {
		// 先把引用的记录删除
		for _, i2 := range list {
			global.DB.Delete(&i2.ChatList)
			logrus.Infof("删除关联对话 %d 条", len(i2.ChatList))
		}
		err = global.DB.Delete(&list).Error
		if err != nil {
			logrus.Error(err)
			response.FailWithMessage("删除会话失败", c)
			return
		}
		logrus.Infof("删除会话 %d 个", len(list))
	}
	response.OkWithMessage(fmt.Sprintf("共删除 %d 个会话", count), c)
}
