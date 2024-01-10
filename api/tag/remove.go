package tag

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/service/common/response"
)

// TagRemoveView 标签删除
// @Tags 标签管理
// @Summary 标签删除
// @Description 标签删除
// @Param data body models.RemoveRequest  true  "查询参数"
// @Param token header string  true  "token"
// @Router /api/tags [delete]
// @Produce json
// @Success 200 {object} response.Response{}
func (TagApi) TagRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}

	var tagList []models.TagModel
	count := global.DB.Find(&tagList, cr.IDList).RowsAffected
	if count == 0 {
		response.FailWithMessage("标签不存在", c)
		return
	}
	// 如果这个标签下有文章，怎么办？
	global.DB.Delete(&tagList)
	response.OkWithMessage(fmt.Sprintf("共删除 %d 个标签", count), c)

}
