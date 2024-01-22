package big_model

import (
	"github.com/gin-gonic/gin"
	"server/config"
	"server/core"
	"server/global"
	"server/service/common/response"
)

// ModelSessionSettingUpdateView 更新会话配置信息
func (BigModelApi) ModelSessionSettingUpdateView(c *gin.Context) {
	var cr config.SessionSetting
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Error(err)
		response.FailWithError(err, &cr, c)
		return
	}

	global.Config.BigModel.SessionSetting = cr
	core.SetYaml()
	response.OkWithMessage("更新成功", c)
	return
}
