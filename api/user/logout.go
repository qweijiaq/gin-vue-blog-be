package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/global"
	logStash "server/plugins/log_stash_v2"
	"server/service"
	"server/service/common/response"
	"server/utils/jwts"
)

// LogoutView 用户注销
// @Tags 用户管理
// @Summary 用户注销
// @Description 用户注销
// @Param token header string  true  "token"
// @Router /api/logout [post]
// @Produce json
// @Success 200 {object} response.Response{}
func (UserApi) LogoutView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	log := logStash.NewAction(c)
	log.SetRequestHeader(c)
	log.SetResponse(c)
	token := c.Request.Header.Get("token")
	err := service.ServiceApp.UserService.Logout(claims, token)

	log.Info(fmt.Sprintf("用户 %s 注销登录", claims.Username))
	if err != nil {
		global.Log.Error(err)
		response.FailWithMessage("注销失败", c)
		return
	}

	response.OkWithMessage("注销成功", c)

}
