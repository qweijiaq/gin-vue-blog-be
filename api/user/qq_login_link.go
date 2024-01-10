package user

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/service/common/response"
)

// QQLoginLinkView 获取qq登录的跳转链接
// @Tags 用户管理
// @Summary 获取qq登录的跳转链接
// @Description 获取qq登录的跳转链接,data就是qq的跳转地址
// @Router /api/qq_login_path [get]
// @Produce json
// @Success 200 {object} response.Response{}
func (UserApi) QQLoginLinkView(c *gin.Context) {
	path := global.Config.QQ.GetPath()
	if path == "" {
		response.FailWithMessage("未配置qq登录地址", c)
		return
	}
	response.OkWithData(path, c)
	return
}
