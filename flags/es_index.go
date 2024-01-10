package flags

import (
	"server/models"
	"server/service/es/indexs"
)

func ESIndex() {
	indexs.CreateIndex(models.ArticleModel{})
	indexs.CreateIndex(models.FullTextModel{})
}
