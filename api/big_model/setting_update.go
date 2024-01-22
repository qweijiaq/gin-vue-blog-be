package big_model

import (
	"github.com/gin-gonic/gin"
	"server/config"
	"server/core"
	"server/global"
	"server/service/common/response"
)

// ModelSettingUpdateView  更新大模型配置
func (BigModelApi) ModelSettingUpdateView(c *gin.Context) {
	var cr config.Setting
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Error(err)
		response.FailWithError(err, &cr, c)
		return
	}

	// 验证 name 是否合法
	var ok bool
	for _, option := range global.Config.BigModel.ModelList {
		if option.Value == cr.Name {
			ok = true
		}
	}
	if !ok {
		response.FailWithMessage("参数「name」错误", c)
	}

	global.Config.BigModel.Setting = cr
	core.SetYaml()
	response.OkWithMessage("更新大模型配置成功", c)
	return
}
