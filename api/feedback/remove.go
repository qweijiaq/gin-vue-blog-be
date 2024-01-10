package feedback

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/service/common/response"
)

// FeedBackRemoveView 删除反馈
// @Tags 反馈管理
// @Summary 删除反馈
// @Description 删除反馈
// @Param data body models.RemoveRequest   true  "表示多个参数"
// @Param token header string  true  "token"
// @Router /api/feedbacks [delete]
// @Produce json
// @Success 200 {object} response.Response{}
func (FeedbackApi) FeedBackRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}

	var list []models.FeedbackModel
	count := global.DB.Find(&list, cr.IDList).RowsAffected
	if count == 0 {
		response.FailWithMessage("内容不存在", c)
		return
	}
	err = global.DB.Delete(&list).Error
	if err != nil {
		response.FailWithMessage("删除反馈内容失败", c)
		return
	}
	response.OkWithMessage(fmt.Sprintf("共删除 %d 条反馈内容", count), c)
}
