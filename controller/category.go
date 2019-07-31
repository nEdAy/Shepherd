package controller

import (
	"github.com/gin-gonic/gin"
)

// @Summary 获取热门搜索关键词
// @Produce  json
// @Success 200 {string} json "{"time": 1561513181, "code": 200, "msg": "成功", "data" : {}}"
// @Failure 400 {string} json "{"time": 1561513181, "code": 400, "msg": "msg"}"
// @Router /v1/category/get-top100" [get]
func GetTop100(c *gin.Context) {
	getFromDataoke(c, Dataoke{
		c.Request.RequestURI,
		2 * 60,
		"https://openapi.dataoke.com/api/category/get-top100",
		map[string]string{
			"version": "v1.0.1",
		}})
}
