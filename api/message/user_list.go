package message

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/service/common/response"
)

type MessageUserListRequest struct {
	models.PageInfo
	NickName string `json:"nickName" form:"nickName"`
}

type MessageUserListResponse struct {
	UserName string `json:"userName"`
	NickName string `json:"nickName"`
	UserID   uint   `json:"userID"`
	Avatar   string `json:"avatar"`
	Count    int    `json:"count"`
}

// MessageUserListView 有消息的用户列表
// @Tags 消息管理
// @Summary 有消息的用户列表
// @Description 有消息的用户列表
// @Router /api/message_users [get]
// @Param token header string  true  "token"
// @Param data query MessageUserListRequest   false  "查询参数"
// @Produce json
// @Success 200 {object} response.Response{data=response.ListResponse[MessageUserListResponse]}
func (MessageApi) MessageUserListView(c *gin.Context) {
	var cr MessageUserListRequest
	c.ShouldBindQuery(&cr)

	var count int64

	global.DB.Model(models.MessageModel{}).Where(models.MessageModel{SendUserNickName: cr.NickName}).
		Group("send_user_id").Count(&count)

	type resType struct {
		SendUserID uint
		Count      int // 发送人的个数2
	}
	offset := (cr.Page - 1) * cr.Limit

	var _list []resType
	global.DB.Model(models.MessageModel{}).Where(models.MessageModel{SendUserNickName: cr.NickName}).
		Group("send_user_id").Limit(cr.Limit).Offset(offset).Select("send_user_id", "count(distinct rev_user_id) as count").Scan(&_list)

	var userMessageMap = map[uint]int{}

	for _, r := range _list {
		userMessageMap[r.SendUserID] = r.Count
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

	response.OkWithList(list, count, c)
}
