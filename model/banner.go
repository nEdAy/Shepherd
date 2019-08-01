package model

import (
	"github.com/jinzhu/gorm"
)

// Banner表
type Banner struct {
	gorm.Model
	Title   string `gorm:"column:title" json:"title"`
	Picture string `gorm:"column:picture" json:"picture"`
	Url     string `gorm:"column:url" json:"url"`
}

// TableName 返回banner表名称
func (Banner) TableName() string {
	return gorm.DefaultTableNameHandler(nil, "banner")
}

// GetAllBanner retrieves all Banner matches certain condition. Returns empty list if no records exist
func GetAllBanner() (banners []*Banner, err error) {
	if err = DB.Order("id desc").Find(&banners).Error; err == nil {
		return banners, nil
	}
	return nil, err
}
