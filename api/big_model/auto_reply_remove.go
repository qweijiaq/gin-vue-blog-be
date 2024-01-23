package big_model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"server/global"
	"server/models"
	"server/service/common/response"
)

// AutoReplyRemoveView 删除自动回复规则
// @Tags 大模型管理
// @Summary 删除自动回复规则
// @Description 删除自动回复规则
// @Param data body models.RemoveRequest  true  "查询参数"
// @Param token header string  true  "token"
// @Router /api/bigModel/auto_reply [delete]
// @Produce json
// @Success 200 {object} response.Response{}
func (BigModelApi) AutoReplyRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}

	var autoReplyList []models.AutoReplyModel
	count := global.DB.Find(&autoReplyList, cr.IDList).RowsAffected
	if count == 0 {
		response.FailWithMessage("自动回复规则不存在", c)
		return
	}

	if len(autoReplyList) > 0 {
		global.DB.Delete(&autoReplyList)
		logrus.Infof("删除自动回复规则 %d 条", len(autoReplyList))
	}
	response.OkWithMessage(fmt.Sprintf("共删除 %d 个自动回复规则", count), c)

}
