package big_model

import (
	"github.com/gin-gonic/gin"
	"server/config"
	"server/global"
	"server/models"
	"server/service/common/response"
	"server/utils/jwts"
)

// ModelSettingView  获取大模型配置
func (BigModelApi) ModelSettingView(c *gin.Context) {
	token := c.GetHeader("token")
	var roleId int
	customClaims, err := jwts.ParseToken(token)
	if err == nil && customClaims != nil {
		roleId = customClaims.Role
	}
	if roleId == models.AdminRole {
		// 管理员
		response.OkWithData(global.Config.BigModel.Setting, c)
		return
	}
	response.OkWithData(config.Setting{
		Enable: global.Config.BigModel.Setting.Enable,
	}, c)
	return
}
