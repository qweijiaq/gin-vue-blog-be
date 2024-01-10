package models

// MenuBannerModel 自定义菜单和背景图的连接表，方便排序
type MenuBannerModel struct {
	MenuID      uint        `gorm:"comment:菜单ID" json:"menu_id"`
	MenuModel   MenuModel   `gorm:"foreignKey:MenuID"`
	BannerID    uint        `gorm:"comment:banner图ID" json:"banner_id"`
	BannerModel BannerModel `gorm:"foreignKey:BannerID"`
	Sort        int         `gorm:"size:10;comment:序号" json:"sort"`
}
