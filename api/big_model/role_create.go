package big_model

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"server/global"
	"server/models"
	"server/service/common/response"
)

type RoleCreateRequest struct {
	Name      string `binding:"required" json:"name"`     // 角色名称
	Enable    bool   `json:"enable"`                      // 是否启用
	Icon      string `json:"icon"`                        // 可以选择系统默认的一些，也可以图片上传
	Abstract  string `binding:"required" json:"abstract"` // 简介
	Scope     int    `json:"scope"`                       // 消耗的积分
	Prologue  string `binding:"required" json:"prologue"` // 开场白
	Prompt    string `binding:"required" json:"prompt"`   // 设定词
	AutoReply bool   `json:"auto_reply"`                  // 自动回复
	TagList   []uint `json:"tagList"`                     // 标签的id列表

}

func (BigModelApi) RoleCreateView(c *gin.Context) {
	var cr RoleCreateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithValidError(err, c)
		return
	}
	// 先找这个标签
	var tags []models.BigModelTagModel
	if len(cr.TagList) == 0 {
		// 这里之所以要判断是因为 Find 如果传入的是空列表，则默认全部查询
		tags = make([]models.BigModelTagModel, 0)
	} else {
		global.DB.Find(&tags, cr.TagList)
		if len(cr.TagList) != len(tags) {
			response.FailWithMessage("标签选择不一致", c)
			return
		}
	}

	// 增加
	var arm models.BigModelRoleModel
	err = global.DB.Take(&arm, "name = ?", cr.Name).Error
	if err == nil {
		response.FailWithMessage("角色名称不能相同", c)
		return
	}
	role := models.BigModelRoleModel{
		Name:      cr.Name,
		Enable:    cr.Enable,
		Icon:      cr.Icon,
		Abstract:  cr.Abstract,
		Scope:     cr.Scope,
		Prologue:  cr.Prologue,
		Prompt:    cr.Prompt,
		AutoReply: cr.AutoReply,
		Tags:      tags, // 会自己关联上
	}
	err = global.DB.Create(&role).Error
	if err != nil {
		logrus.Errorf("角色添加失败 err：%s, 原始数据内容 %#v", err, cr)
		response.FailWithMessage("角色添加失败", c)
		return
	}
	response.OkWithMessage("角色添加成功", c)
	return
}
