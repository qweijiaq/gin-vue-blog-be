package big_model

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/service/common/response"
	"server/utils/jwts"
)

type SessionUserUpdateNameRequest struct {
	SessionID uint   `json:"sessionID" binding:"required"` // 会话id
	Name      string `json:"name"`
}

// SessionUpdateNameView 用户修改会话名称
func (BigModelApi) SessionUpdateNameView(c *gin.Context) {
	var cr SessionUserUpdateNameRequest
	// control request
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

	if session.UserID != claims.UserID {
		response.FailWithMessage("会话鉴权失败", c)
		return
	}
	// 修改会话名称
	err = global.DB.Model(&session).Updates(models.BigModelSessionModel{Name: cr.Name}).Error
	if err != nil {
		response.FailWithMessage("会话名称修改失败", c)
		return
	}
	response.OkWithMessage("会话名称修改成功", c)

}
