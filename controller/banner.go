package controller

import (
	"Shepherd/model"
	"Shepherd/pkg/response"
	"github.com/gin-gonic/gin"
)

// @Summary 获取Banners
func GetBanners(c *gin.Context) {
	banners, err := model.GetAllBanner()
	if err != nil {
		response.ErrorWithMsg(c, err.Error())
	} else {
		response.JsonWithData(c, banners)
	}
}
