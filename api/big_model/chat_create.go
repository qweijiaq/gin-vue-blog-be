package big_model

import (
	broadscope_bailian "github.com/aliyun/alibabacloud-bailian-go-sdk/client"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"io"
	"server/global"
	"server/models"
	"server/service/big_model"
	"server/service/common/response"
	"server/utils/jwts"
)

type ChatCreateRequest struct {
	SessionID uint   `form:"sessionID" json:"sessionID" binding:"required"` // 会话id
	Content   string `form:"content" json:"content" binding:"required"`     // 对话内容
}

// ChatCreateView 当前用户创建对话
func (BigModelApi) ChatCreateView(c *gin.Context) {
	// 先认证
	token := c.Query("token")
	claims, err := jwts.ParseToken(token)
	if err != nil {
		response.FailWithMessageSSE("认证失败", c)
		return
	}

	var cr ChatCreateRequest
	err = c.ShouldBindQuery(&cr)
	if err != nil {
		response.FailWithMessageSSE("参数错误", c)
		return
	}

	// 找会话
	var session models.BigModelSessionModel
	err = global.DB.Take(&session, cr.SessionID).Error
	if err != nil {
		response.FailWithMessageSSE("会话不存在", c)
		return
	}
	// 验证这个会话是不是自己创建的
	if session.UserID != claims.UserID {
		response.FailWithMessageSSE("对话鉴权错误", c)
		return
	}

	// 判断这个用户能不能创建对话
	var user models.UserModel
	err = global.DB.Take(&user, claims.UserID).Error
	if err != nil {
		response.FailWithMessageSSE("用户信息错误", c)
		return
	}

	scope := global.Config.BigModel.SessionSetting.ChatScope

	if user.Scope-scope <= 0 {
		response.FailWithMessageSSE("积分不足，无法创建对话", c)
		return
	}

	reply, err := big_model.Send(cr.SessionID, cr.Content)
	if err != nil {
		response.FailWithMessageSSE(err.Error(), c)
		return
	}

	var botContent string
	c.Stream(func(w io.Writer) bool {
		if r, ok := <-reply.(chan *broadscope_bailian.CompletionResponse); ok {
			response.OkWithDataSSE(r.GetData().GetText(), c)
			botContent = r.GetData().GetText()
			return true
		}
		return false
	})

	var chatModel = models.BigModelChatModel{
		SessionID:  cr.SessionID,
		Status:     true,
		Content:    cr.Content,
		BotContent: botContent,
		RoleID:     session.RoleID,
		UserID:     claims.UserID,
	}
	err = global.DB.Create(&chatModel).Error
	if err != nil {
		response.FailWithMessageSSE("对话创建失败", c)
		return
	}

	// 扣用户的积分
	global.DB.Model(&user).Update("scope", gorm.Expr("scope - ?", scope))
	response.OkWithSSE(chatModel.ID, "ok", c)
}
