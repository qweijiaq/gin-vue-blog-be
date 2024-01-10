package images

import (
	"github.com/gin-gonic/gin"
	"io/fs"
	"os"
	"server/global"
	"server/service"
	"server/service/common/response"
	"server/service/image"
	"server/utils/jwts"
)

// ImageUploadView 上传多个图片，返回图片的 URL
// @Tags 图片管理
// @Summary 批量上传图片
// @Description 上传多个图片，返回图片的 URL
// @Param token header string  true  "token"
// @Accept multipart/form-data
// @Param images formData file true "文件上传"
// @Router /api/images [post]
// @Produce json
// @Success 200 {object} response.Response{}
func (ImagesApi) ImageUploadView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	if claims.Role == 3 {
		response.FailWithMessage("游客不支持上传图片", c)
		return
	}
	// 上传多个图片
	form, err := c.MultipartForm()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	fileList, ok := form.File["images"]
	if !ok {
		response.FailWithMessage("文件不存在", c)
		return
	}

	// 判断路径是否存在, 不存在就创建
	basePath := global.Config.Upload.Path
	_, err = os.ReadDir(basePath)
	if err != nil {
		// 递归创建
		err = os.MkdirAll(basePath, fs.ModePerm)
		if err != nil {
			global.Log.Error(err)
		}
	}

	var resList []image.FileUploadResponse

	for _, file := range fileList {

		// 上传文件
		serviceRes := service.ServiceApp.ImageService.ImageUploadService(file)
		if !serviceRes.IsSuccess {
			resList = append(resList, serviceRes)
			continue
		}
		if !global.Config.QiNiu.Enable {
			// 本地保存
			err = c.SaveUploadedFile(file, serviceRes.FileName)
			if err != nil {
				global.Log.Error(err)
				serviceRes.Msg = err.Error()
				serviceRes.IsSuccess = false
				resList = append(resList, serviceRes)
				continue
			}
		}
		resList = append(resList, serviceRes)
	}

	response.OkWithData(resList, c)

}
