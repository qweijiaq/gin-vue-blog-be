package big_model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"regexp"
	"server/global"
	"server/models"
	"server/service/common/response"
)

type AutoReplyUpdateRequest struct {
	Id           uint   `json:"id"`
	Name         string `json:"name" binding:"required"`
	Mode         int    `json:"mode" binding:"required,oneof=1 2 3 4"`
	Rule         string `json:"rule" binding:"required"`
	ReplyContent string `json:"reply_content" binding:"required"`
}

// AutoReplyUpdateView 增加和修改自动回复
func (BigModelApi) AutoReplyUpdateView(c *gin.Context) {
	var cr models.AutoReplyModel
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithValidError(err, c)
		return
	}

	// 如果是正则规则，校验正则表达式格式
	if cr.Mode == 4 {
		_, err := regexp.Compile(cr.Rule)
		if err != nil {
			response.FailWithMessage(fmt.Sprintf("正则表达式格式错误 %s", err.Error()), c)
			return
		}
	}
	if cr.ID == 0 {
		// 	增加
		var arm models.AutoReplyModel
		err = global.DB.Take(&arm, "name = ?", cr.Name).Error
		if err == nil {
			response.FailWithMessage("规则名称不能相同", c)
			return
		}
		err = global.DB.Create(&models.AutoReplyModel{
			Name:         cr.Name,
			Mode:         cr.Mode,
			Rule:         cr.Rule,
			ReplyContent: cr.ReplyContent,
		}).Error
		if err != nil {
			logrus.Errorf("规则添加失败 err: %s, 原始数据内容 %#v", err, cr)
			response.FailWithMessage("规则添加失败", c)
			return
		}
		response.FailWithMessage("规则添加成功", c)
		return
	}

	// 更新
	var arm models.AutoReplyModel
	err = global.DB.Take(&arm, cr.ID).Error
	if err != nil {
		response.FailWithMessage("记录不存在", c)
		return
	}
	// name 不能再重复了
	var arm1 models.AutoReplyModel
	err = global.DB.Take(&arm1, "name = ? and id <> ?", cr.Name, cr.ID).Error
	if err == nil {
		response.FailWithMessage("规则名称不能和已有规则相重复", c)
		return
	}
	err = global.DB.Model(&arm).Updates(map[string]any{
		"name":          cr.Name,
		"mode":          cr.Mode,
		"rule":          cr.Rule,
		"reply_content": cr.ReplyContent,
	}).Error
	if err != nil {
		logrus.Errorf("规则更新失败 err: %s, 原始数据内容 %#v", err, cr)
		response.FailWithMessage("规则更新失败", c)
		return
	}
	response.OkWithMessage("规则更新成功", c)
	return
}
