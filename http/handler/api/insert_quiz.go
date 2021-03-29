package api

import (
	"ecnu.space/tmp-loser/model"
	"github.com/gin-gonic/gin"
)

func InsertQuiz(c *gin.Context){
	qq := model.QuizQuestion{}
	err := c.ShouldBind(&qq)
	if err != nil {
		
	}
}