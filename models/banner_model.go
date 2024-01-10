package models

import (
	"gorm.io/gorm"
	"os"
	"server/global"
	"server/models/ctype"
)

type BannerModel struct {
	MODEL
	Path        string            `gorm:"comment:图片路径" json:"path"`                                      // 图片路径
	Hash        string            `gorm:"comment:图片的hash值" json:"hash"`                                  // 图片的 hash 值，用于判断重复图片
	Name        string            `gorm:"size:38;comment:图片名称" json:"name"`                              // 图片名称
	ImageType   ctype.ImageType   `gorm:"default:1;comment:图片的类型，本地还是七牛,1本地，2七牛，默认是1" json:"image_type"` // 图片的类型， 本地还是七牛
	MenusBanner []MenuBannerModel `gorm:"foreignKey:BannerID" json:"-"`
}

func (b *BannerModel) BeforeDelete(tx *gorm.DB) (err error) {
	if b.ImageType == ctype.Local {
		// 本地图片，删除，还要删除本地的存储
		err = os.Remove(b.Path[1:])
		if err != nil {
			global.Log.Error(err)
			return nil
		}
	}
	return nil
}
