package big_model

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"server/global"
	"server/models"
	"server/service/common/response"
)

type TagUpdateRequest struct {
	ID    uint   `json:"id"`                       // 更新使用
	Title string `json:"title" binding:"required"` // 名称
	Color string `json:"color" binding:"required"` // 颜色
}

// TagUpdateView 标签新增和更新
func (BigModelApi) TagUpdateView(c *gin.Context) {
	var cr TagUpdateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithValidError(err, c)
		return
	}

	if cr.ID == 0 {
		// 增加
		var arm models.BigModelTagModel
		err = global.DB.Take(&arm, "title = ?", cr.Title).Error
		if err == nil {
			response.FailWithMessage("标签名称不能相同", c)
			return
		}
		err = global.DB.Create(&models.BigModelTagModel{
			Title: cr.Title,
			Color: cr.Color,
		}).Error
		if err != nil {
			logrus.Errorf("数据添加失败 err：%s, 原始数据内容 %#v", err, cr)
			response.FailWithMessage("角色标签添加失败", c)
			return
		}
		response.OkWithMessage("角色标签添加成功", c)
		return
	}
	var arm models.BigModelTagModel
	err = global.DB.Take(&arm, cr.ID).Error
	if err != nil {
		response.FailWithMessage("记录不存在", c)
		return
	}
	// name 不能和已有的重复
	var arm1 models.BigModelTagModel
	err = global.DB.Take(&arm1, "title = ? and id <> ?", cr.Title, cr.ID).Error
	if err == nil {
		response.FailWithMessage("和已有的标签名称重复", c)
		return
	}
	err = global.DB.Model(&arm).Updates(map[string]any{
		"title": cr.Title,
		"color": cr.Color,
	}).Error
	if err != nil {
		logrus.Errorf("数据更新失败 err：%s, 原始数据内容 %#v", err, cr)
		response.FailWithMessage("角色标签更新失败", c)
		return
	}
	response.OkWithMessage("角色标签更新成功", c)
}
