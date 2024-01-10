package models

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"server/models/ctype"
)

// UserModel 用户表
type UserModel struct {
	MODEL
	NickName   string           `gorm:"size:36;comment:昵称" json:"nick_name,select(c|info)"`                    // 昵称
	UserName   string           `gorm:"size:36;comment:用户名" json:"user_name"`                                  // 用户名
	Password   string           `gorm:"size:128;comment:密码" json:"-"`                                          // 密码
	Avatar     string           `gorm:"size:256;comment:头像" json:"avatar,select(c)"`                           // 头像
	Email      string           `gorm:"size:128;comment:邮箱" json:"email,select(info)"`                         // 邮箱
	Tel        string           `gorm:"size:18;comment:手机号" json:"tel"`                                        // 手机号
	Addr       string           `gorm:"size:64;comment:地址" json:"addr,select(c|info)"`                         // 地址
	Token      string           `gorm:"size:64;comment:其他平台的唯一id" json:"token"`                                // 其他平台的唯一id
	IP         string           `gorm:"size:20;comment:ip" json:"ip,select(c)"`                                // ip地址
	Role       ctype.Role       `gorm:"size:4;default:1;comment:权限，1管理员，2普通用户，3游客" json:"role,select(info)"`   // 权限  1 管理员  2 普通用户  3 游客
	SignStatus ctype.SignStatus `gorm:"type=smallint(6);comment:注册来源，1qq，3邮箱" json:"sign_status,select(info)"` // 注册来源
	Integral   int              `gorm:"default:0;comment:我的积分" json:"integral,select(info)"`                   // 我的积分
	Sign       string           `gorm:"size:128;comment:我的签名" json:"sign,select(info)"`                        // 我的签名
	Link       string           `gorm:"size:128;comment:我的链接地址" json:"link,select(info)"`                      // 我的链接地址
}

func (u *UserModel) BeforeDelete(tx *gorm.DB) (err error) {

	var loginDataList []LoginDataModel
	err = tx.Find(&loginDataList, "user_id = ?", u.ID).Updates(map[string]any{
		"user_id": nil,
	}).Error
	if err != nil {
		logrus.Error(err)
		return err
	}
	logrus.Infof("删除关联 登录数据 %d 条", len(loginDataList))

	var collects []UserCollectModel
	err = tx.Find(&collects, "user_id = ?", u.ID).Delete(&collects).Error
	if err != nil {
		logrus.Error(err)
		return err
	}
	logrus.Infof("删除关联 收藏文章数据 %d 条", len(collects))

	var messageList []MessageModel
	err = tx.Find(&messageList, "send_user_id = ? or rev_user_id = ?", u.ID, u.ID).Delete(&messageList).Error
	if err != nil {
		logrus.Error(err)
		return err
	}

	logrus.Infof("删除关联 用户消息 %d 条", len(messageList))

	var commentList []CommentModel
	err = tx.Find(&commentList, "user_id = ?", u.ID).Updates(map[string]any{
		"user_id": nil,
	}).Error
	if len(commentList) > 0 {
		err = tx.Delete(&commentList).Error
		if err != nil {
			logrus.Error(err)
			return err
		}
	}
	logrus.Infof("删除关联 评论数据 %d 条", len(commentList))
	return
}
