package advert

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/service/common/response"
)

// AdvertRemoveView 批量删除广告
// @Tags 广告管理
// @Summary 删除广告
// @Description 批量删除广告
// @Param data body models.RemoveRequest    true  "广告 ID 列表"
// @Param token header string  true  "token"
// @Router /api/adverts [delete]
// @Produce json
// @Success 200 {object} response.Response{}
func (AdvertApi) AdvertRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}

	var advertList []models.AdvertModel
	count := global.DB.Find(&advertList, cr.IDList).RowsAffected
	if count == 0 {
		response.FailWithMessage("广告不存在", c)
		return
	}
	global.DB.Delete(&advertList)
	response.OkWithMessage(fmt.Sprintf("共删除 %d 个广告", count), c)

}
