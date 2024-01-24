package big_model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"server/global"
	"server/models"
	"server/service/common/response"
)

// ChatRemoveView 管理员删除对话
func (BigModelApi) ChatRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithValidError(err, c)
		return
	}

	var list []models.BigModelChatModel
	count := global.DB.Find(&list, cr.IDList).RowsAffected
	if count == 0 {
		response.FailWithMessage("记录不存在", c)
		return
	}

	if len(list) > 0 {
		// 先把引用的记录删除
		err = global.DB.Delete(&list).Error
		if err != nil {
			logrus.Error(err)
			response.FailWithMessage("删除对话失败", c)
			return
		}
		logrus.Infof("删除对话 %d 个", len(list))
	}
	response.OkWithMessage(fmt.Sprintf("共删除 %d 个对话", count), c)
}
