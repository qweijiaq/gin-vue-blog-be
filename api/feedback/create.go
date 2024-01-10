package feedback

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/service/common/response"
)

type FeedBackCreate struct {
	Email   string `json:"email" binding:"required,email"`
	Content string `json:"content" binding:"required"`
}

// FeedBackCreateView 发布反馈
// @Tags 反馈管理
// @Summary 发布反馈
// @Description 发布反馈
// @Param data body FeedBackCreate true  "参数"
// @Router /api/feedbacks [post]
// @Produce json
// @Success 200 {object} response.Response{}
func (FeedbackApi) FeedBackCreateView(c *gin.Context) {
	var cr FeedBackCreate
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithError(err, &cr, c)
		return
	}

	var model models.FeedbackModel
	err = global.DB.Take(&model, "email = ? and content = ?", cr.Email, cr.Content).Error
	if err == nil {
		response.FailWithMessage("存在相同留言", c)
		return
	}
	global.DB.Create(&models.FeedbackModel{
		Email:   cr.Email,
		Content: cr.Content,
	})
	response.OkWithMessage("反馈成功", c)
}
