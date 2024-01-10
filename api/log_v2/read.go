package log_v2

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	logStash "server/plugins/log_stash_v2"
	"server/service/common/response"
)

// LogReadView 日志读取
// @Tags 日志管理V2
// @Summary 日志读取
// @Description 日志读取
// @Description 1. 前端判断这个日志的读取状态，未读就去请求这个接口，让这个日志变成已读的
// @Description 2. 如果是已读状态，就不需要调这个接口了
// @Param data query models.IDRequest true "参数"
// @Param token header string true "token"
// @Router /api/logs/v2/read [get]
// @Produce json
// @Success 200 {object} response.Response{}
func (LogApi) LogReadView(c *gin.Context) {
	var cr models.IDRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	var log logStash.LogModel
	err = global.DB.Take(&log, cr.ID).Error
	if err != nil {
		response.FailWithMessage("日志不存在", c)
		return
	}
	if log.ReadStatus {
		response.OkWithMessage("日志读取成功", c)
		return
	}
	global.DB.Model(&log).Update("readStatus", true)
	response.OkWithMessage("日志读取成功", c)
	return
}
