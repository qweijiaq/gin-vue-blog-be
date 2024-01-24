package big_model

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"server/global"
	"server/models"
	"server/service/common/response"
	"server/utils/jwts"
)

type ChatCreateRequest struct {
	SessionID uint   `json:"sessionID" binding:"required"` // 会话id
	Content   string `json:"content" binding:"required"`   // 对话内容
}

// ChatCreateView 当前用户创建对话
func (BigModelApi) ChatCreateView(c *gin.Context) {
	var cr ChatCreateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithValidError(err, c)
		return
	}

	// 找会话
	var session models.BigModelSessionModel
	err = global.DB.Take(&session, cr.SessionID).Error
	if err != nil {
		response.FailWithMessage("会话不存在", c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	// 验证这个会话是不是自己创建的
	if session.UserID != claims.UserID {
		response.FailWithMessage("对话鉴权错误", c)
		return
	}

	// 判断这个用户能不能创建对话
	var user models.UserModel
	err = global.DB.Take(&user, claims.UserID).Error
	if err != nil {
		response.FailWithMessage("用户信息错误", c)
		return
	}

	scope := global.Config.BigModel.SessionSetting.ChatScope

	if user.Scope-scope <= 0 {
		response.FailWithMessage("积分不足，无法创建对话", c)
		return
	}

	err = global.DB.Create(&models.BigModelChatModel{
		SessionID:  cr.SessionID,
		Status:     true,
		Content:    cr.Content,
		BotContent: "你好",
		RoleID:     session.RoleID,
		UserID:     claims.UserID,
	}).Error
	if err != nil {
		response.FailWithMessage("对话创建失败", c)
		return
	}
	// 扣用户的积分
	global.DB.Model(&user).Update("scope", gorm.Expr("scope - ?", scope))
	response.Ok("你好", "对话创建成功", c)
}
