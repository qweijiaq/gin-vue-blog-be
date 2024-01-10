package chat

import (
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
	"server/models"
	"server/service/common"
	"server/service/common/response"
)

// ChatListView 群聊聊天记录
// @Tags 聊天管理
// @Summary 群聊聊天记录
// @Description 群聊聊天记录
// @Param data query models.PageInfo   false  "表示多个参数"
// @Router /api/chat_groups_records [get]
// @Produce json
// @Success 200 {object} response.Response{data=response.ListResponse[models.ChatModel]}
func (ChatApi) ChatListView(c *gin.Context) {
	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}

	cr.Sort = "created_at desc"
	list, count, _ := common.ComList(models.ChatModel{IsGroup: true}, common.Option{
		PageInfo: cr,
	})

	data := filter.Omit("list", list)
	_list, _ := data.(filter.Filter)
	if string(_list.MustMarshalJSON()) == "{}" {
		list = make([]models.ChatModel, 0)
		response.OkWithList(list, count, c)
		return
	}
	response.OkWithList(data, count, c)

}
