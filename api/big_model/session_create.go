package big_model

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"server/global"
	"server/models"
	"server/service/common/response"
	"server/utils/jwts"
)

type SessionCreateRequest struct {
	RoleID uint `json:"roleID" binding:"required"` // 角色id
}

// SessionCreateView 当前用户创建会话
func (BigModelApi) SessionCreateView(c *gin.Context) {
	// controller request
	var cr SessionCreateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithValidError(err, c)
		return
	}

	// 找角色
	var role models.BigModelRoleModel
	err = global.DB.Take(&role, cr.RoleID).Error
	if err != nil {
		response.FailWithMessage("角色不存在", c)
		return
	}

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	// 判断这个用户能不能创建会话
	var user models.UserModel
	err = global.DB.Take(&user, claims.UserID).Error
	if err != nil {
		response.FailWithMessage("用户信息错误", c)
		return
	}

	scope := global.Config.BigModel.SessionSetting.SessionScope

	if user.Scope-scope <= 0 {
		response.FailWithMessage("积分不足，无法创建会话", c)
		return
	}

	// 名字默认就是新的会话
	// 如果用户创建了一个新的会话，但是没有聊天，那就不能创建
	// 找这个用户相关的 AI 角色，有没有空的对话记录 > 1
	var sessionList []models.BigModelSessionModel
	global.DB.Preload("ChatList").Find(&sessionList, "user_id = ? and role_id = ?", claims.UserID, cr.RoleID)
	var ok bool
	var sessionID uint
	for _, model := range sessionList {
		if len(model.ChatList) <= 0 {
			ok = true
			sessionID = model.ID
		}
	}
	if ok {
		response.Ok(sessionID, "已存在新的会话", c)
		return
	}

	var session = models.BigModelSessionModel{
		UserID: claims.UserID,
		RoleID: cr.RoleID,
		Name:   "新的会话",
	}
	global.DB.Create(&session)

	// 扣用户的积分
	global.DB.Model(&user).Update("scope", gorm.Expr("scope - ?", scope))
	response.Ok(session.ID, "会话创建成功", c)

}
