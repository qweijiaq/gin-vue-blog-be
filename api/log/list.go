package log

import (
	"github.com/gin-gonic/gin"
	"server/models"
	"server/plugins/log_stash"
	"server/service/common"
	"server/service/common/response"
)

type LogRequest struct {
	models.PageInfo
	Level log_stash.Level `form:"level"`
}

// LogListView 日志列表
func (LogApi) LogListView(c *gin.Context) {
	var cr LogRequest
	c.ShouldBindQuery(&cr)
	list, count, _ := common.ComList(log_stash.LogStashModel{Level: cr.Level}, common.Option{
		PageInfo: cr.PageInfo,
		Debug:    true,
		Likes:    []string{"ip", "addr"},
	})
	response.OkWithList(list, count, c)
	return
}
