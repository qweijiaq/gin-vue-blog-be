package big_model

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"server/global"
	"server/models"
	"server/service/common/response"
	"server/utils/jwts"
)

type UserScopeRequest struct {
	Status bool `json:"status"`
}

// UserScopeView 用户是否可以领取积分
func (BigModelApi) UserScopeView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	userId := claims.UserID

	var cr UserScopeRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	// 查这个用户今天能不能领取这个积分
	var userScopeModel models.UserScopeModel
	err = global.DB.Take(&userScopeModel, "user_id = ? and to_days(created_at)=to_days(now())", userId).Error
	if err == nil {
		// 查到了
		response.FailWithMessage("今日已领取积分啦", c)
		return
	}

	var user models.UserModel
	err = global.DB.Take(&user, userId).Error
	if err != nil {
		response.FailWithMessage("用户不存在", c)
		return
	}
	if cr.Status == false {
		response.OkWithMessage("", c)
		return
	}
	// 给用户加积分
	scope := global.Config.BigModel.SessionSetting.DayScope
	global.DB.Model(&user).Update("scope", gorm.Expr("scope + ?", scope))
	global.DB.Create(&models.UserScopeModel{
		UserId: userId,
		Scope:  scope,
		Status: cr.Status,
	})
	response.OkWithMessage("积分领取成功", c)
	return
}
