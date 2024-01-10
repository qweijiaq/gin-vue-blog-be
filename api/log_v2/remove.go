package log_v2

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/global"
	logStash "server/plugins/log_stash_v2"
	"server/service/common/response"
	"time"
)

type LogRemoveRequest struct {
	IDList    []uint `json:"id_list"`   // 可以传id列表删除
	StartTime string `json:"startTime"` // 年月日格式的开始时间
	EndTime   string `json:"endTime"`   // 年月日格式的结束时间
	UserID    uint   `json:"userID"`    // 根据用户删除日志
	IP        string `json:"ip"`        // 根据用户ip删除
}

// LogRemoveView 删除日志
// @Tags 日志管理V2
// @Summary 删除日志
// @Description 删除日志
// @Param data body LogRemoveRequest true "参数"
// @Param token header string true "token"
// @Router /api/logs/v2 [delete]
// @Produce json
// @Success 200 {object} response.Response{}
func (LogApi) LogRemoveView(c *gin.Context) {
	var cr LogRemoveRequest
	log := logStash.NewAction(c)
	log.SetRequest(c)
	log.SetResponse(c)
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	var logs []logStash.LogModel
	if len(cr.IDList) > 0 {
		log.SetItemInfo("IDList", cr.IDList)
		global.DB.Find(&logs, cr.IDList).Delete(&logs)
	} else if cr.UserID != 0 {
		global.DB.Find(&logs, "user_id = ?", cr.UserID).Delete(&logs)
	} else if cr.IP != "" {
		global.DB.Find(&logs, "ip = ?", cr.IP).Delete(&logs)
	} else if cr.StartTime != "" && cr.EndTime != "" {
		_, startTimeErr := time.Parse("2006-01-02", cr.StartTime)
		_, endTimeErr := time.Parse("2006-01-02", cr.EndTime)
		if startTimeErr != nil {
			response.FailWithMessage("开始时间格式错误", c)
			return
		}
		if endTimeErr != nil {
			response.FailWithMessage("结束时间格式错误", c)
			return
		}
		global.DB.Find(&logs, "created_at > date(?) and created_at < date(?)", cr.StartTime, cr.EndTime).Delete(&logs)
	}

	log.SetItemInfo("共删除日志", len(logs))
	log.Info("删除日志成功")

	response.OkWithMessage(fmt.Sprintf("共删除 %d 条日志", len(logs)), c)
}
