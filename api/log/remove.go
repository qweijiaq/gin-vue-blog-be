package log

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/plugins/log_stash"
	"server/service/common/response"
)

// LogRemoveListView 删除日志
func (LogApi) LogRemoveListView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}

	var list []log_stash.LogStashModel
	count := global.DB.Find(&list, cr.IDList).RowsAffected
	if count == 0 {
		response.FailWithMessage("日志不存在", c)
		return
	}
	global.DB.Delete(&list)
	response.OkWithMessage(fmt.Sprintf("共删除 %d 个日志", count), c)

}
