package big_model

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/service/common/response"
)

// ModelUsableListView 获取可用大模型列表
func (BigModelApi) ModelUsableListView(c *gin.Context) {
	response.OkWithData(global.Config.BigModel.ModelList, c)
	return
}
