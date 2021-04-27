package middleware

import (
	"ecnu.space/tmp-loser/utils"
	"github.com/gin-gonic/gin"
)

func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetString("authorization")
		if authorization != "1" && authorization != "2" {
			utils.HandleNotAdminErr(c, "非管理员")
			c.Abort()
		}
		// 处理请求
		c.Next() //  处理请求
	}
}
