package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/nEdAy/Shepherd/model"
	"github.com/nEdAy/Shepherd/pkg/response"
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
