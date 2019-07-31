package controller

import "github.com/gin-gonic/gin"

// @Summary 获取品牌库信息
// @Produce  json
// @Success 200 {string} json "{"time": 1561513181, "code": 200, "msg": "成功", "data" : {}}"
// @Failure 400 {string} json "{"time": 1561513181, "code": 400, "msg": "msg"}"
// @Router /api/tb-service/get-brand-list" [get]
func GetBrandList(c *gin.Context) {
	getFromDataoke(c, Dataoke{
		c.Request.RequestURI,
		2 * 60,
		"https://openapi.dataoke.com/api/tb-service/get-brand-list",
		map[string]string{
			"version": "v1.0.2",
		}})
}
