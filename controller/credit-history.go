package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nEdAy/Shepherd/model"
	"github.com/nEdAy/Shepherd/pkg/jwt"
	"github.com/nEdAy/Shepherd/pkg/response"
	"math/rand"
	"time"
)

func GetCreditHistoriesByUserId(c *gin.Context) {
	userId, _ := c.Get(jwt.KeyUserId)
	creditHistories, err := model.GetCreditHistoriesByUserId(userId.(uint))
	if err != nil {
		response.ErrorWithMsg(c, err.Error())
	} else {
		response.JsonWithData(c, creditHistories)
	}
}

func Shake(c *gin.Context) {
	userId, _ := c.Get(jwt.KeyUserId)
	// 随机获取积分
	rand.Seed(time.Now().UnixNano())
	change := -1
	x := rand.Intn(100)
	if x < 40 { // 40%
		change = 10
	} else if x < 70 { // 30%
		change = 20
	} else if x < 90 { // 20%
		change = 30
	} else { // 10%
		change = 40
	}
	message := fmt.Sprintf("每日摇一摇获得%d积分", change)
	// 更新积分并创建积分记录
	creditHistory, err := model.ModifyCreditHistory(userId.(uint), change, message)
	if err != nil {
		response.ErrorWithMsg(c, err.Error())
	} else {
		response.JsonWithData(c, creditHistory)
	}
}
