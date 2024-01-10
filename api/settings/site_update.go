package settings

import (
	"github.com/gin-gonic/gin"
	"server/config"
	"server/core"
	"server/global"
	"server/service/common/response"
)

// SettingsSiteUpdateView 编辑网站信息
// @Tags 系统管理
// @Summary 编辑网站信息
// @Description 编辑网站信息
// @Param data body config.SiteInfo true "编辑网站信息的参数"
// @Param token header string  true  "token"
// @Router /api/settings/site [put]
// @Produce json
// @Success 200 {object} response.Response{data=config.SiteInfo}
func (SettingsApi) SettingsSiteUpdateView(c *gin.Context) {
	var info config.SiteInfo
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}
	global.Config.SiteInfo = info
	err = core.SetYaml()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("网站信息更新成功", c)
}
