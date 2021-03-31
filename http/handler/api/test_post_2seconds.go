package api

import (
	"time"

	"github.com/gin-gonic/gin"
)

func TestPost2Seconds(c *gin.Context) {
	time.Sleep(time.Second * 2)
	c.JSON(200, struct {
		Haha string
	}{
		Haha: "你好",
	})
}
