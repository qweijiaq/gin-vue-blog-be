package news

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"io"
	"server/service/common/response"
	"server/service/redis"
	"server/utils/requests"
	"time"
)

type params struct {
	ID   string `json:"id"`
	Size int    `json:"size"`
}

type header struct {
	Signaturekey string `form:"signaturekey" structs:"signaturekey"`
	Version      string `form:"version" structs:"version"`
	UserAgent    string `form:"User-Agent" structs:"User-Agent"`
}

type NewResponse struct {
	Code int             `json:"code"`
	Data []redis.NewData `json:"data"`
	Msg  string          `json:"msg"`
}

const newAPI = "https://api.codelife.cc/api/top/list"
const timeout = 2 * time.Second

// NewListView 新闻列表
// @Tags 新闻管理
// @Summary 新闻列表
// @Description 新闻列表
// @Param data body params true "新闻参数"
// @Param version header string true "version"
// @Param User-Agent header string true "User-Agent"
// @Param signaturekey header string true "signaturekey"
// @Router /api/news [post]
// @Produce json
// @Success 200 {object} response.Response{data=[]redis.NewData}
func (NewsApi) NewListView(c *gin.Context) {
	var cr params
	var headers header
	err := c.ShouldBindJSON(&cr)
	err = c.ShouldBindHeader(&headers)
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}
	if cr.Size == 0 {
		cr.Size = 1
	}

	key := fmt.Sprintf("%s-%d", cr.ID, cr.Size)
	newsData, _ := redis.GetNews(key)
	if len(newsData) != 0 {
		response.OkWithData(newsData, c)
		return
	}

	httpResponse, err := requests.Post(newAPI, cr, structs.Map(headers), timeout)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	var res NewResponse
	byteData, err := io.ReadAll(httpResponse.Body)
	err = json.Unmarshal(byteData, &res)
	if err != nil {
		fmt.Println(string(byteData))
		response.FailWithMessage(err.Error(), c)
		return
	}
	if res.Code != 200 {
		response.FailWithMessage(res.Msg, c)
		return
	}
	response.OkWithData(res.Data, c)
	redis.SetNews(key, res.Data)
	return
}
