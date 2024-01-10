package article

import (
	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday"
	"math/rand"
	"server/global"
	"server/models"
	"server/service/common/response"
	"server/service/es"
	"server/utils/jwts"
	"strings"
	"time"
)

type ArticleRequest struct {
	Title    string   `json:"title" binding:"required" msg:"文章标题必填"`   // 文章标题
	Abstract string   `json:"abstract"`                                // 文章简介
	Content  string   `json:"content" binding:"required" msg:"文章内容必填"` // 文章内容
	Category string   `json:"category"`                                // 文章分类
	Source   string   `json:"source"`                                  // 文章来源
	Link     string   `json:"link"`                                    // 原文链接
	BannerID uint     `json:"banner_id"`                               // 文章封面 ID
	Tags     []string `json:"tags"`                                    // 文章标签
}

// ArticleCreateView 发布文章
// @Tags 文章管理
// @Summary 发布文章
// @Description 发布文章
// @Param data body ArticleRequest   true  "表示多个参数"
// @Param token header string  true  "token"
// @Router /api/articles [post]
// @Produce json
// @Success 200 {object} response.Response{}
func (ArticleApi) ArticleCreateView(c *gin.Context) {
	var cr ArticleRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithError(err, &cr, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	userID := claims.UserID
	userNickName := claims.NickName
	// 校验 content  xss

	// 处理 content
	unsafe := blackfriday.MarkdownCommon([]byte(cr.Content))
	// 是不是有 script 标签
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string(unsafe)))
	//fmt.Println(doc.Text())
	nodes := doc.Find("script").Nodes
	if len(nodes) > 0 {
		// 有 script 标签
		doc.Find("script").Remove()
		converter := md.NewConverter("", true, nil)
		html, _ := doc.Html()
		markdown, _ := converter.ConvertString(html)
		cr.Content = markdown
	}
	if cr.Abstract == "" {
		// 汉字的截取不一样
		abs := []rune(doc.Text())
		// 将 content 转为 html，并且过滤 xss，以及获取中文内容
		if len(abs) > 100 {
			cr.Abstract = string(abs[:100])
		} else {
			cr.Abstract = string(abs)
		}
	}

	// 不传 banner_id, 后台就随机去选择一张
	if cr.BannerID == 0 {
		var bannerIDList []uint
		global.DB.Model(models.BannerModel{}).Select("id").Scan(&bannerIDList)
		if len(bannerIDList) == 0 {
			response.FailWithMessage("没有banner数据", c)
			return
		}
		rand.Seed(time.Now().UnixNano())
		cr.BannerID = bannerIDList[rand.Intn(len(bannerIDList))]
	}

	// 查 banner_id 下的 banner_url
	var bannerUrl string
	err = global.DB.Model(models.BannerModel{}).Where("id = ?", cr.BannerID).Select("path").Scan(&bannerUrl).Error
	if err != nil {
		response.FailWithMessage("banner不存在", c)
		return
	}

	// 查用户头像
	var avatar string
	err = global.DB.Model(models.UserModel{}).Where("id = ?", userID).Select("avatar").Scan(&avatar).Error
	if err != nil {
		response.FailWithMessage("用户不存在", c)
		return
	}
	now := time.Now().Format("2006-01-02 15:04:05")

	article := models.ArticleModel{
		CreatedAt:    now,
		UpdatedAt:    now,
		Title:        cr.Title,
		Keyword:      cr.Title,
		Abstract:     cr.Abstract,
		Content:      cr.Content,
		UserID:       userID,
		UserNickName: userNickName,
		UserAvatar:   avatar,
		Category:     cr.Category,
		Source:       cr.Source,
		Link:         cr.Link,
		BannerID:     cr.BannerID,
		BannerUrl:    bannerUrl,
		Tags:         cr.Tags,
	}
	// 应该去判断文章标题是否存在
	if article.ISExistData() {
		response.FailWithMessage("文章已存在", c)
		return
	}
	err = article.Create()
	if err != nil {
		global.Log.Error(err)
		response.FailWithMessage(err.Error(), c)
		return
	}

	go es.AsyncArticleByFullText(article.ID, article.Title, article.Content)

	response.OkWithMessage("文章发布成功", c)

}
