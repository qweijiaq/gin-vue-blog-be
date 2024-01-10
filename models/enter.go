package models

import "time"

type MODEL struct {
	ID        uint      `gorm:"primaryKey;comment:id" json:"id,select($any)" structs:"-"` // 主键ID
	CreatedAt time.Time `gorm:"comment:创建时间" json:"created_at,select($any)" structs:"-"`  // 创建时间
	UpdatedAt time.Time `gorm:"comment:更新时间" json:"-" structs:"-"`                        // 更新时间
}
type RemoveRequest struct {
	IDList []uint `json:"id_list"`
}

type IDRequest struct {
	ID uint `json:"id" form:"id" uri:"id"`
}

type ESIDRequest struct {
	ID string `json:"id" form:"id" uri:"id"`
}
type ESIDListRequest struct {
	IDList []string `json:"id_list" binding:"required"`
}

type PageInfo struct {
	Page  int    `form:"page"`
	Key   string `form:"key"`
	Limit int    `form:"limit"`
	Sort  string `form:"sort"`
}
