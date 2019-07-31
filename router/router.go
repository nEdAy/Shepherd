package router

import (
	"Shepherd/controller"
	_ "Shepherd/docs"
	"Shepherd/middleware"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var Router *gin.Engine

func Setup() {
	Router = gin.Default()
	Router.Use(gin.Logger())
	Router.Use(gin.Recovery())
	Router.Static("/assets", "./assets")
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// Ping test
	Router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	// v1
	v1 := Router.Group("/v1")
	{
		// 注册用户 / 用户登录(密码) / 用户登录（短信验证码）
		v1.POST("/registerOrLogin", controller.RegisterOrLogin)
		// 用户相关API
		user := v1.Group("/user", middleware.JWT())
		{
			// 获取用户
			user.GET("/", controller.GetUser)
			// 删除用户
			// user.DELETE("/:id", controller.DelUser)
		}
		// Banner相关API
		banner := v1.Group("/banner")
		{
			// 获取Banners
			banner.GET("/", controller.GetBanners)
		}
		// 商品相关API
		goods := v1.Group("/goods")
		{
			// 获取各大榜单
			goods.GET("/get-ranking-list", controller.GetRankingList)
			// 获取9.9精选
			goods.GET("/nine/op-goods-list", controller.GetNineOpGoodsList)
			// 获取猜你喜欢
			goods.GET("/list-similer-goods-by-open", controller.ListSimilerGoodsByOpen)
			// 获取超级搜索
			goods.GET("/list-super-goods", controller.ListSuperGoods)
			// 获取大淘客搜索
			goods.GET("/get-dtk-search-goods", controller.GetDtkSearchGoods)

		}
		// 分类相关API
		category := v1.Group("/category")
		{
			// 获取热搜记录
			category.GET("/get-top100", controller.GetTop100)
		}
	}
}
