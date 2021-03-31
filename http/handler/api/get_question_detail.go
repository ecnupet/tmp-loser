package api

import (
	"log"
	"strconv"

	"ecnu.space/tmp-loser/model"
	"ecnu.space/tmp-loser/store"
	"ecnu.space/tmp-loser/utils"
	"github.com/gin-gonic/gin"
)

func GetQuestionDetail(c *gin.Context) {
	questionId := c.Query("questionId")
	log.Println(questionId)
	qid, err := strconv.Atoi(questionId)
	log.Println(qid)
	if err != nil {
		utils.HandleGetDBErr(c, "GetQuestionDetail Atoi fail")
		return
	}
	qs, err := store.GetDB().QuestionRW.GetQuestionById(uint32(qid))
	if err != nil {
		utils.HandleGetDBErr(c, "GetQuestionDetail GetQuestionById Fail")
		return
	}
	if len(qs) < 0 {
		utils.HandleGetDBErr(c, "GetQuestionDetail questionId doesn't exist")
		return
	}
	q := qs[0]

	utils.HandleGetSuccess(c, model.QuestionStringTime{
		QuestionID:  q.QuestionID,
		Description: q.Description,
		Type:        q.Type,
		Options:     q.Options,
		Duration:    q.Duration,
	})
}
