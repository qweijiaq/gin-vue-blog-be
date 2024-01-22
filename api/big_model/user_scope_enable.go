package big_model

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/service/common/response"
	"server/utils/jwts"
)

type UserScopeEnableResponse struct {
	Enable bool `json:"enable"` // 用户能不能领取
	Scope  int  `json:"scope"`  // 能领取多少积分
}

// UserScopeEnableView 用户是否可以领取积分
func (BigModelApi) UserScopeEnableView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	userId := claims.UserID

	var res UserScopeEnableResponse

	// 查这个用户今天能不能领取这个积分
	var userScopeModel models.UserScopeModel
	err := global.DB.Take(&userScopeModel, "user_id = ? and to_days(created_at)=to_days(now())", userId).Error
	if err == nil {
		// 查到了
		response.OkWithData(res, c)
		return
	}
	res.Enable = true
	res.Scope = global.Config.BigModel.SessionSetting.DayScope
	response.OkWithData(res, c)
	return
}
