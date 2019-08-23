package model

import (
	"github.com/jinzhu/gorm"
)

type Banner struct {
	gorm.Model
	Index   uint
	Title   string
	Picture string
	Url     string
}

func GetAllBanner() (banners []*Banner, err error) {
	if err = db.Order("banner.index ASC").Find(&banners).Error; err == nil {
		return banners, nil
	}
	return nil, err
}
