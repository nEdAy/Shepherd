package model

import (
	"github.com/jinzhu/gorm"
)

type Banner struct {
	gorm.Model
	Index   uint   `json:"index"`
	Title   string `json:"title"`
	Picture string `json:"picture"`
	Url     string `json:"url"`
}

func GetAllBanner() (banners []*Banner, err error) {
	if err = db.Order("banner.index ASC").Find(&banners).Error; err == nil {
		return banners, nil
	}
	return nil, err
}
