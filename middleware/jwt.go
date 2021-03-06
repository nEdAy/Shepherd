package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/nEdAy/Shepherd/pkg/jwt"
	"github.com/nEdAy/Shepherd/pkg/response"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			response.ErrorWithMsg(c, "Token不存在")
			c.Abort()
			return
		} else {
			claims, err := jwt.ParseToken(token)
			if err != nil {
				response.ErrorWithMsg(c, err.Error())
				c.Abort()
				return
			} else {
				userId := claims.UserId
				if userId < 0 {
					response.ErrorWithMsg(c, "Token异常，用户不存在")
					c.Abort()
					return
				}
				c.Set(jwt.KeyUserId, claims.UserId)
				c.Next()
			}
		}
	}
}
