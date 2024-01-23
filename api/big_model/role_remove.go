package big_model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"server/global"
	"server/models"
	"server/service/common/response"
)

// RoleRemoveView 删除角色
// @Tags 大模型管理
// @Summary 删除角色
// @Description 删除角色
// @Param data body models.RemoveRequest  true  "查询参数"
// @Param token header string  true  "token"
// @Router /api/bigModel/roles [delete]
// @Produce json
// @Success 200 {object} response.Response{}
func (BigModelApi) RoleRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}

	var roleList []models.BigModelRoleModel
	count := global.DB.Preload("Tags").Find(&roleList, cr.IDList).RowsAffected
	if count == 0 {
		response.FailWithMessage("角色不存在", c)
		return
	}

	if len(roleList) > 0 {
		// 先把引用的记录删除
		for _, i2 := range roleList {
			global.DB.Model(&i2).Association("Tags").Delete(i2.Tags)
		}
		err = global.DB.Delete(&roleList).Error
		if err != nil {
			logrus.Error(err)
			response.FailWithMessage("删除角色失败", c)
			return
		}
		logrus.Infof("删除角色标签 %d 个", len(roleList))
	}
	response.OkWithMessage(fmt.Sprintf("共删除 %d 个角色", count), c)

}
