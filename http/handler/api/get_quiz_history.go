package api

import (
	"log"
	"strconv"

	"ecnu.space/tmp-loser/model"
	"ecnu.space/tmp-loser/store"
	"ecnu.space/tmp-loser/utils"
	"github.com/gin-gonic/gin"
)

const (
	timeFormat = "2006-01-02 15:04:05"
)

func GetQuizHistory(c *gin.Context) {
	userName := c.Query("userName")
	page := c.Query("page")
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		utils.HandleGetErr(c, "GetQuizHistory page Atoi err")
		return
	}
	pageSize := c.Query("pageSize")
	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil {
		utils.HandleGetErr(c, "GetQuizHistory pageSize Atoi err")
		return
	}

	quizInfos := []model.QuizHistoryResult{}

	qids, err := store.GetDB().CommitHistoryRW.GetQuizIDByUserNameAndPageNoAndNum(userName, uint32(pageInt), uint32(pageSizeInt))
	if err != nil {
		utils.HandleGetErr(c, "GetQuizHistory no quizID err")
		return
	}
	for _, qid := range qids {
		types := []uint32{}
		costTime := uint32(0)
		chs, err := store.GetDB().CommitHistoryRW.GetAllCommitsHistoryByUserNameAndQuizID(userName, qid)
		if err != nil {
			utils.HandleGetDBErr(c, "GetQuizHistory GetAllCommitsHistoryByUserNameAndQuizID err: "+err.Error())
		}

		if len(chs) == 0 {
			utils.HandleGetDBErr(c, "GetQuizHistory no commit history in quiz")
			return
		}
		startTime := chs[0].CreatedAt.Format(timeFormat)
		point := uint32(0)
		for _, ch := range chs {
			log.Println(ch)
			costTime += ch.Spend
			if ch.Correct == 1 {
				point += 10
			}
			qs, err := store.GetDB().QuestionRW.GetQuestionById(ch.QuestionID)
			if err != nil {
				utils.HandleGetDBErr(c, "GetQuizHistory GetQuestionById err :"+err.Error())
				return
			}
			for _, q := range qs {
				if notIn(types, q.Type) {
					types = append(types, q.Type)
				}
			}
		}
		quizInfos = append(quizInfos, model.QuizHistoryResult{
			QuizID:    qid,
			Types:     types,
			StartTime: startTime,
			Point:     point,
			CostTime:  costTime,
		})
	}
	utils.HandleGetSuccess(c, quizInfos)

}

func notIn(arr []uint32, n uint32) bool {
	for _, x := range arr {
		if n == x {
			return false
		}
	}
	return true
}
