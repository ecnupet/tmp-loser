package api

import (
	"log"

	"ecnu.space/tmp-loser/model"
	"ecnu.space/tmp-loser/store"
	"ecnu.space/tmp-loser/utils"
	"github.com/gin-gonic/gin"
)

func InsertQuestion(c *gin.Context) {
	q := model.Question{}
	err := c.ShouldBind(&q)
	if err != nil {
		log.Println("InsertQuestion bind  err:", err)
		utils.HandlePostQuizQuestionErr(c, "InsertQuestion bind error: "+err.Error())
		return
	}

	err = store.GetDB().QuestionRW.Insert(&q)
	if err != nil {
		log.Println("InsertQuestion err: "+err.Error())
		utils.HandlePostDBErr(c, "题目插入失败： " + err.Error())
		return 
	}
	utils.HandlePostSuccess(c, struct{
		Info   string `json:"info"`
	}{
		Info: "插入题目成功",
	})
}
