package big_model

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"reflect"
	"server/global"
	"server/models"
	"server/service/common/response"
)

type RoleUpdateRequest struct {
	ID        uint    `binding:"required" json:"id" structs:"-"`
	Name      *string `json:"name" structs:"name"`            // 角色名称
	Enable    *bool   `json:"enable" structs:"enable"`        // 是否启用
	Icon      *string `json:"icon" structs:"icon"`            // 可以选择系统默认的一些，也可以图片上传
	Abstract  *string `json:"abstract" structs:"abstract"`    // 简介
	Scope     *int    `json:"scope" structs:"scope"`          // 消耗的积分
	Prologue  *string `json:"prologue" structs:"prologue"`    // 开场白
	Prompt    *string `json:"prompt" structs:"prompt"`        // 设定词
	AutoReply *bool   `json:"autoReply" structs:"auto_reply"` // 自动回复
	TagList   *[]uint `json:"tagList" structs:"-"`            // 标签的id列表
}

func (BigModelApi) RoleUpdateView(c *gin.Context) {
	var cr RoleUpdateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithValidError(err, c)
		return
	}
	var arm models.BigModelRoleModel
	err = global.DB.Preload("Tags").Take(&arm, cr.ID).Error
	if err != nil {
		response.FailWithMessage("记录不存在", c)
		return
	}

	if cr.TagList != nil {
		// 先找这个标签
		var tags []models.BigModelTagModel
		if len(*cr.TagList) == 0 {
			// 这里之所以要判断是因为 Find 如果传入的是空列表，则默认全部查询
			tags = make([]models.BigModelTagModel, 0)
		} else {
			global.DB.Find(&tags, cr.TagList)
			if len(*cr.TagList) != len(tags) {
				response.FailWithMessage("标签选择不一致", c)
				return
			}
		}
		// 把之前的标签删掉，再新增
		global.DB.Model(&arm).Association("Tags").Replace(tags) // 替换
	}

	maps := structs.Map(cr)
	var modelMap = map[string]any{}

	for key, mp := range maps {
		// 把 nil 过滤掉
		val := reflect.ValueOf(mp)
		if val.Kind() == reflect.Ptr && val.IsNil() {
			continue
		}
		modelMap[key] = val.Elem().Interface()

		if key == "name" {
			// 校验新的名字是不是重复了
			var m1 models.BigModelRoleModel
			err = global.DB.Take(&m1, "name = ? and id <> ?", val.Elem().String(), cr.ID).Error
			if err == nil {
				response.FailWithMessage("角色名称和已有的角色名称重复", c)
				return
			}
		}
	}
	if len(modelMap) == 0 {
		response.OkWithMessage("角色数据暂无更新", c)
		return
	}
	err = global.DB.Model(&arm).Updates(modelMap).Error
	if err != nil {
		logrus.Error(err)
		response.FailWithMessage("角色更新失败", c)
		return
	}

	response.OkWithMessage("角色更新成功", c)
}
