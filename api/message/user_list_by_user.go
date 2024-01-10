package message

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/service/common/response"
)

type MessageUserListByUserRequest struct {
	UserID uint `json:"userID" form:"userID" binding:"required"`
}

// MessageUserListByUserView 某个用户的聊天列表
// @Tags 消息管理
// @Summary 某个用户的聊天列表
// @Description 某个用户的聊天列表
// @Router /api/message_users/user [get]
// @Param token header string  true  "token"
// @Param data query MessageUserListByUserRequest   false  "查询参数"
// @Produce json
// @Success 200 {object} response.Response{data=response.ListResponse[MessageUserListResponse]}
func (MessageApi) MessageUserListByUserView(c *gin.Context) {
	var cr MessageUserListByUserRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	type resType struct {
		SendUserID uint
		RevUserID  uint
		Count      int
	}

	var _list []resType
	global.DB.Model(models.MessageModel{}).Where("send_user_id = ? or rev_user_id = ?", cr.UserID, cr.UserID).
		Group("send_user_id").
		Group("rev_user_id").Select("send_user_id", "rev_user_id", "count(id) as count").Scan(&_list)

	var userMessageMap = map[uint]int{}

	for _, r := range _list {
		sendVal, ok1 := userMessageMap[r.SendUserID]
		if !ok1 && cr.UserID != r.SendUserID {
			userMessageMap[r.SendUserID] = r.Count
		} else {
			if cr.UserID != r.SendUserID {
				userMessageMap[r.SendUserID] = r.Count + sendVal
			}
		}
		revVal, ok2 := userMessageMap[r.RevUserID]
		if !ok2 && cr.UserID != r.RevUserID {
			userMessageMap[r.RevUserID] = r.Count
		} else {
			if cr.UserID != r.RevUserID {
				userMessageMap[r.RevUserID] = r.Count + revVal
			}
		}
	}
	var userIDList []uint
	for uid, _ := range userMessageMap {
		userIDList = append(userIDList, uid)
	}
	var userList []models.UserModel
	global.DB.Find(&userList, userIDList)
	var userMap = map[uint]models.UserModel{}
	for _, model := range userList {
		userMap[model.ID] = model
	}

	var list = make([]MessageUserListResponse, 0)
	for uid, count := range userMessageMap {
		user := userMap[uid]
		list = append(list, MessageUserListResponse{
			UserName: user.UserName,
			NickName: user.NickName,
			UserID:   user.ID,
			Avatar:   user.Avatar,
			Count:    count,
		})
	}

	response.OkWithList(list, int64(len(list)), c)
}
