package user

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/ctype"
	log_stash "server/plugins/log_stash_v2"
	"server/service/common/response"
	"server/utils"
	"server/utils/jwts"
	"server/utils/pwd"
)

type EmailLoginRequest struct {
	UserName string `json:"user_name" binding:"required" msg:"请输入用户名"`
	Password string `json:"password" binding:"required" msg:"请输入密码"`
}

// EmailLoginView 邮箱登录，返回token，用户信息需要从token中解码
// @Tags 用户管理
// @Summary 邮箱登录
// @Description 邮箱登录，返回token，用户信息需要从token中解码
// @Param data body EmailLoginRequest  true  "查询参数"
// @Router /api/email_login [post]
// @Produce json
// @Success 200 {object} response.Response{}
func (UserApi) EmailLoginView(c *gin.Context) {
	var cr EmailLoginRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithError(err, &cr, c)
		return
	}

	var userModel models.UserModel
	err = global.DB.Take(&userModel, "user_name = ? or email = ?", cr.UserName, cr.UserName).Error
	if err != nil {
		// 没找到
		global.Log.Warn("用户名不存在")
		log_stash.NewFailLogin("用户名不存在", cr.UserName, cr.Password, c)
		response.FailWithMessage("用户名或密码错误", c)
		return
	}
	// 校验密码
	isCheck := pwd.CheckPwd(userModel.Password, cr.Password)
	if !isCheck {
		global.Log.Warn("用户名密码错误")
		log_stash.NewFailLogin("用户名密码错误", cr.UserName, cr.Password, c)
		response.FailWithMessage("用户名或密码错误", c)
		return
	}
	// 登录成功，生成token
	token, err := jwts.GenToken(jwts.JwtPayLoad{
		NickName: userModel.NickName,
		Role:     int(userModel.Role),
		UserID:   userModel.ID,
		Username: userModel.UserName,
	})
	if err != nil {
		global.Log.Error(err)
		response.FailWithMessage("token生成失败", c)
		return
	}
	ip, addr := utils.GetAddrByGin(c)
	c.Request.Header.Set("token", token)
	log_stash.NewSuccessLogin(c)

	global.DB.Create(&models.LoginDataModel{
		UserID:    userModel.ID,
		IP:        ip,
		NickName:  userModel.NickName,
		Token:     token,
		Device:    "",
		Addr:      addr,
		LoginType: ctype.SignEmail,
	})

	response.OkWithData(token, c)

}
