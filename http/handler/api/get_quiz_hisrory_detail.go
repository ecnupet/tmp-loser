package api

import (
	"log"
	"strconv"
	"time"

	"ecnu.space/tmp-loser/model"
	"ecnu.space/tmp-loser/store"
	"ecnu.space/tmp-loser/utils"
	"github.com/gin-gonic/gin"
)

func GetQuizHistoryDetail(c *gin.Context) {
	userNameAny, e := c.Get("user_name")
	if userNameAny == nil && !e {
		utils.HandleGetDBErr(c, "userName quizId 缺一不可")
		return

	}
	userName := userNameAny.(string)
	quizId := c.Query("quizId")
	if e && userName == "" || quizId == "" {
		utils.HandleGetDBErr(c, "userName quizId 缺一不可")
		return
	}
	quizIdInt, err := strconv.Atoi(quizId)
	if err != nil {
		utils.HandleGetErr(c, "GetQuizDetail quizId Atoi fail")
		return
	}
	chs, err := store.GetDB().CommitHistoryRW.GetAllCommitsHistoryByUserNameAndQuizID(userName, uint32(quizIdInt))
	if err != nil {
		utils.HandleGetDBErr(c, "GetQuizDetail err : "+err.Error())
		return
	}
	if len(chs) == 0 {
		utils.HandleGetErr(c, "quiz have no commit history")
		return
	}
	startTime := chs[0].CreatedAt.Add(time.Hour * 8).Format(timeFormat)
	costTime := uint32(0)
	qrs := []model.QuizHistoryDetailResultQuestion{}
	for _, ch := range chs {
		costTime += ch.Spend
		// get question info
		qinfos, err := store.GetDB().QuestionRW.GetQuestionById(ch.QuestionID)
		if len(qinfos) == 0 {
			utils.HandleGetDBErr(c, "GetQuizDetail questionId doesn't exist "+err.Error())
			return
		}
		qinfo := qinfos[0]
		if err != nil {
			utils.HandleGetDBErr(c, "GetQuizDetail get question info err: "+err.Error())
			return
		}
		qr := model.QuizHistoryDetailResultQuestion{
			QuestionID:  ch.QuestionID,
			Description: qinfo.Description,
			Type:        qinfo.Type,
			Options:     qinfo.Options,
			Duration:    qinfo.Duration,
			Answer:      qinfo.Answer,
			Choice:      ch.Choose,
			Spend:       ch.Spend,
		}
		log.Println("choice为： ", ch.Choose)
		if ch.Choose == "" {
			log.Println("choice 是空串")
			qr.Choice = nil
		}
		qrs = append(qrs, qr)
	}
	r := model.QuizHistoryDetailResult{
		StartTime: startTime,
		CostTime:  costTime,
		Results:   qrs,
	}
	utils.HandleGetSuccess(c, r)
}
