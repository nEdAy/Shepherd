package controller

import (
	"Shepherd/pkg/response"
	"github.com/gin-gonic/gin"
)

type rankingList struct {
	RankType string `form:"rankType" binding:"required"`
	Cid      string `form:"cid" `
}

func GetRankingList(c *gin.Context) {
	var rankingList rankingList
	if err := c.ShouldBindQuery(&rankingList); err != nil {
		response.ErrorWithMsg(c, err.Error())
		return
	}
	getFromDataoke(c, Dataoke{
		c.Request.RequestURI,
		2 * 60,
		"https://openapi.dataoke.com/api/goods/get-ranking-list",
		map[string]string{
			"version":  "v1.0.2",
			"rankType": rankingList.RankType,
			"cid":      rankingList.Cid,
		}})
}

type nineOpGoodsList struct {
	PageSize string `form:"pageSize" binding:"required"`
	PageId   string `form:"pageId" binding:"required"`
	Cid      string `form:"cid" binding:"required"`
}

func GetNineOpGoodsList(c *gin.Context) {
	var nineOpGoodsList nineOpGoodsList
	if err := c.ShouldBindQuery(&nineOpGoodsList); err != nil {
		response.ErrorWithMsg(c, err.Error())
		return
	}
	getFromDataoke(c, Dataoke{
		c.Request.RequestURI,
		2 * 60,
		"https://openapi.dataoke.com/api/goods/nine/op-goods-list",
		map[string]string{
			"version":  "v1.0.1",
			"pageSize": nineOpGoodsList.PageSize,
			"pageId":   nineOpGoodsList.PageId,
			"cid":      nineOpGoodsList.Cid,
		}})
}

type listSimilerGoodsByOpen struct {
	Id   string `form:"id" binding:"required"`
	Size string `form:"size" `
}

func ListSimilerGoodsByOpen(c *gin.Context) {
	var listSimilerGoodsByOpen listSimilerGoodsByOpen
	if err := c.ShouldBindQuery(&listSimilerGoodsByOpen); err != nil {
		response.ErrorWithMsg(c, err.Error())
		return
	}
	getFromDataoke(c, Dataoke{
		c.Request.RequestURI,
		2 * 60,
		"https://openapi.dataoke.com/api/goods/list-similer-goods-by-open",
		map[string]string{
			"version": "v1.0.1",
			"id":      listSimilerGoodsByOpen.Id,
			"size":    listSimilerGoodsByOpen.Size,
		}})
}

type listSuperGoods struct {
	Type     string `form:"type"`
	KeyWords string `form:"keyWords" binding:"required"`
	Tmall    string `form:"tmall"`
	Haitao   string `form:"haitao"`
	Sort     string `form:"sort"`
}

func ListSuperGoods(c *gin.Context) {
	var listSuperGoods listSuperGoods
	if err := c.ShouldBindQuery(&listSuperGoods); err != nil {
		response.ErrorWithMsg(c, err.Error())
		return
	}
	getFromDataoke(c, Dataoke{
		c.Request.RequestURI,
		2 * 60,
		"https://openapi.dataoke.com/api/goods/list-super-goods",
		map[string]string{
			"version":  "v1.0.1",
			"type":     listSuperGoods.Type,
			"keyWords": listSuperGoods.KeyWords,
			"tmall":    listSuperGoods.Tmall,
			"haitao":   listSuperGoods.Haitao,
			"sort":     listSuperGoods.Sort,
		}})
}

type dtkSearchGoods struct {
	PageSize string `form:"pageSize" binding:"required"`
	PageId   string `form:"pageId" binding:"required"`
	KeyWords string `form:"keyWords" binding:"required"`
}

func GetDtkSearchGoods(c *gin.Context) {
	var dtkSearchGoods dtkSearchGoods
	if err := c.ShouldBindQuery(&dtkSearchGoods); err != nil {
		response.ErrorWithMsg(c, err.Error())
		return
	}
	getFromDataoke(c, Dataoke{
		c.Request.RequestURI,
		2 * 60,
		"https://openapi.dataoke.com/api/goods/get-dtk-search-goods",
		map[string]string{
			"version":  "v2.0.0",
			"pageSize": dtkSearchGoods.PageSize,
			"pageId":   dtkSearchGoods.PageId,
			"keyWords": dtkSearchGoods.KeyWords,
		}})
}
