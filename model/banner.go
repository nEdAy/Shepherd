package model

import (
	"github.com/jinzhu/gorm"
)

type Banner struct {
	gorm.Model
	Title   string
	Picture string
	Url     string
}

func GetAllBanner() (banners []*Banner, err error) {
	if err = db.Order("index asc").Find(&banners).Error; err == nil {
		return banners, nil
	}
	return nil, err
}
