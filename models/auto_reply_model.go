package models

import (
	"regexp"
	"server/global"
	"strings"
)

// AutoReplyModel 自动回复表
type AutoReplyModel struct {
	MODEL
	Name         string `gorm:"size:32" json:"name"`            // 规则名称
	Mode         int    `json:"mode"`                           // 匹配模式 1-精确匹配 2-模糊匹配 3-前缀匹配 4-正则匹配
	Rule         string `gorm:"size:64" json:"rule"`            // 匹配规则
	ReplyContent string `gorm:"size:1024" json:"reply_content"` // 回复内容
}

// HitAutoReplyRule 命中自动回复规则
func (AutoReplyModel) HitAutoReplyRule(content string) *AutoReplyModel {
	var list []AutoReplyModel
	global.DB.Find(&list)
	for _, model := range list {
		switch model.Mode {
		case 1:
			if model.Rule == content {
				return &model
			}
		case 2:
			if strings.Contains(content, model.Rule) {
				return &model
			}
		case 3:
			if strings.HasPrefix(content, model.Rule) {
				return &model
			}
		case 4:
			// 接口那里校验了正则表达式格式，这里不会出错
			regex, _ := regexp.Compile(model.Rule)
			if regex.MatchString(content) {
				return &model
			}
		}
	}
	return nil
}
