package role

import (
	"github.com/gin-gonic/gin"
	"server/service/common/response"
)

type OptionResponse struct {
	Label string `json:"label"`
	Value int    `json:"value"`
}

func (RoleApi) RoleIDListView(c *gin.Context) {
	response.OkWithData([]OptionResponse{
		{"管理员", 1},
		{"普通用户", 2},
		{"游客", 3},
	}, c)
}
