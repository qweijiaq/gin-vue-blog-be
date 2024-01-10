package tag

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/service/common/response"
)

// TagUpdateView 更新标签
// @Tags 标签管理
// @Summary 更新标签
// @Description 更新标签
// @Param data body TagRequest  true "查询参数"
// @Param token header string  true  "token"
// @Param id path int  true  "id"
// @Router /api/tags/{id} [put]
// @Produce json
// @Success 200 {object} response.Response{}
func (TagApi) TagUpdateView(c *gin.Context) {

	id := c.Param("id")
	var cr TagRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithError(err, &cr, c)
		return
	}
	var tag models.TagModel
	err = global.DB.Take(&tag, id).Error
	if err != nil {
		response.FailWithMessage("标签不存在", c)
		return
	}
	// 结构体转map的第三方包
	maps := structs.Map(&cr)
	err = global.DB.Model(&tag).Updates(maps).Error

	if err != nil {
		global.Log.Error(err)
		response.FailWithMessage("修改标签失败", c)
		return
	}

	response.OkWithMessage("修改标签成功", c)
}
