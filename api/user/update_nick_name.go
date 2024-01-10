package user

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/ctype"
	"server/service/common/response"
	"server/utils/jwts"
	"strings"
)

type UserUpdateNicknameRequest struct {
	NickName string `json:"nick_name" structs:"nick_name"`
	Sign     string `json:"sign" structs:"sign"`
	Link     string `json:"link" structs:"link"`
	Avatar   string `json:"avatar" structs:"avatar"`
}

// UserUpdateNickName 修改当前登录人的昵称，签名，链接
// @Tags 用户管理
// @Summary 修改当前登录人的昵称，签名，链接
// @Description 修改当前登录人的昵称，签名，链接
// @Router /api/user_info [put]
// @Param token header string  true  "token"
// @Param data body UserUpdateNicknameRequest  true  "昵称，签名，链接"
// @Produce json
// @Success 200 {object} response.Response{}
func (UserApi) UserUpdateNickName(c *gin.Context) {
	var cr UserUpdateNicknameRequest
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}
	var newMaps = map[string]interface{}{}
	maps := structs.Map(cr)
	for key, v := range maps {
		if val, ok := v.(string); ok && strings.TrimSpace(val) != "" {
			newMaps[key] = val
		}
	}

	var userModel models.UserModel
	err = global.DB.Debug().Take(&userModel, claims.UserID).Error
	if err != nil {
		response.FailWithMessage("用户不存在", c)
		return
	}

	// 如果改的是头像，则判断一下用户的注册来源
	_, ok := newMaps["avatar"]
	if ok && userModel.SignStatus != ctype.SignEmail {
		delete(newMaps, "avatar")
	}

	err = global.DB.Model(&userModel).Updates(newMaps).Error
	if err != nil {
		global.Log.Error(err)
		response.FailWithMessage("修改用户信息失败", c)
		return
	}
	response.OkWithMessage("修改个人信息成功", c)

}
