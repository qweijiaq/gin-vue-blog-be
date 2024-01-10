package log_stash

import (
	"github.com/gin-gonic/gin"
	"server/global"
)

// NewSuccessLogin 登录成功的日志
func NewSuccessLogin(c *gin.Context) {
	token := c.Request.Header.Get("token")
	jwyPayLoad := parseToken(token)
	saveLoginLog("登录成功", "--", jwyPayLoad.UserID, jwyPayLoad.UserName, true, c)
}

// NewFailLogin 登录失败的日志
func NewFailLogin(title, userName, pwd string, c *gin.Context) {
	saveLoginLog(title, pwd, 0, userName, false, c)
}

func saveLoginLog(title string, content string, userID uint, userName string, status bool, c *gin.Context) {
	ip := c.ClientIP()
	addr := getAddr(ip)
	global.DB.Create(&LogModel{
		IP:       ip,
		Addr:     addr,
		Title:    title,
		Content:  content,
		UserID:   userID,
		UserName: userName,
		Status:   status,
		Type:     LoginType,
	})
}
