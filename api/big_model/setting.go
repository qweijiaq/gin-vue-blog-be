package big_model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"path"
	"server/config"
	"server/global"
	"server/models"
	"server/service/common/response"
	"server/utils/jwts"
)

const docsPath = "uploads/docs"

type ModelSetting struct {
	config.Setting
	Help string `json:"help"`
}

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
		ms := ModelSetting{
			Setting: global.Config.BigModel.Setting,
		}
		if ms.Name != "" {
			filePath := path.Join(docsPath, fmt.Sprintf("%s.md", ms.Name))
			byteData, err := os.ReadFile(filePath)
			if err == nil {
				ms.Help = string(byteData)
			}
		}
		response.OkWithData(ms, c)
		return
	}
	response.OkWithData(ModelSetting{
		Setting: config.Setting{
			Enable: global.Config.BigModel.Setting.Enable,
			Title:  global.Config.BigModel.Setting.Title,
			Slogan: global.Config.BigModel.Setting.Slogan,
		},
	}, c)
	return
}
