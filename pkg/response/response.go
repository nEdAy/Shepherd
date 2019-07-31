package response

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

func ErrorWithMsg(c *gin.Context, msg string) {
	log.Error().Msg(msg)
	c.JSON(http.StatusBadRequest, gin.H{"time": time.Now().Unix(), "code": http.StatusBadRequest, "msg": msg})
}

func JsonWithData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"time": time.Now().Unix(), "code": 0, "msg": "成功", "data": data})
}
