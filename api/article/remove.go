package article

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"server/global"
	"server/models"
	"server/service/common/response"
	"server/service/es"
)

type IDListRequest struct {
	IDList []string `json:"id_list"`
}

// ArticleRemoveView 删除文章
// @Tags 文章管理
// @Summary 删除文章
// @Description 批量删除文章
// @Param data body IDListRequest   true  "表示多个参数"
// @Param token header string  true  "token"
// @Router /api/articles [delete]
// @Produce json
// @Success 200 {object} response.Response{}
func (ArticleApi) ArticleRemoveView(c *gin.Context) {
	var cr IDListRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Error(err)
		response.FailWithCode(response.ArgumentError, c)
		return
	}

	// 如果文章删除了，用户收藏这篇文章怎么办？
	// 1. 顺带把与这个文章关联的收藏也删除了
	// 2. 用户收藏表，新增一个字段，表示文章是否删除，用户可以删除这个收藏记录，但是找不到文章去改收藏数
	bulkService := global.ESClient.Bulk().Index(models.ArticleModel{}.Index()).Refresh("true")
	for _, id := range cr.IDList {
		req := elastic.NewBulkDeleteRequest().Id(id)
		bulkService.Add(req)
		go es.DeleteFullTextByArticleID(id)
	}
	result, err := bulkService.Do(context.Background())
	if err != nil {
		global.Log.Error(err)
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage(fmt.Sprintf("成功删除 %d 篇文章", len(result.Succeeded())), c)
	return

}
