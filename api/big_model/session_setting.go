package big_model

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/service/common/response"
)

// ModelSessionSettingView 获取会话配置信息
func (BigModelApi) ModelSessionSettingView(c *gin.Context) {
	response.OkWithData(global.Config.BigModel.SessionSetting, c)
	return
}
