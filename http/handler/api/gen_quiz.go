package api

import (
	"log"

	"ecnu.space/tmp-loser/db"
	"ecnu.space/tmp-loser/model"
	"ecnu.space/tmp-loser/store"
	"ecnu.space/tmp-loser/utils"
	"github.com/gin-gonic/gin"
)

const (
	requiredQuestionNum = 1
)

// 鉴权配置前，先假设提供user_name参数
func GenQuiz(c *gin.Context) {
	tt := model.NewQuizParams{}
	err := c.ShouldBind(&tt)
	if err != nil {
		log.Println("GenQuiz  err:", err)
		utils.HandlePostQuizQuestionErr(c, err.Error())
		return
	}
	userName := tt.UserName
	ts := tt.Types
	ratesMap := make(map[uint32]float32)
	for _, t := range ts {
		questions, err := store.GetDB().QuestionRW.GetQuestionByType(t)
		if err != nil {
			log.Println("GenQuiz GetQuestionByType error: ", err)
			utils.HandleGetDBErr(c, err.Error())
			return
		}
		for _, question := range questions {
			ratesMap[question.QuestionID] = GetQuestionCorrectRateByUser(userName, question.QuestionID)
		}
	}

	// // 从小到大
	// sort.Float64Slice(rates).Sort()
	if len(ratesMap) < requiredQuestionNum {
		log.Println("GenQuiz shortage in question num")
		utils.HandleGetNumErr(c, "GenQuiz shortage in question num")
		return
	}
	quizQuestionRst := []uint32{}
	for k, _ := range ratesMap {
		quizQuestionRst = append(quizQuestionRst, k)
	}
	quizIDs, err := store.GetDB().CommitHistoryRW.GetQuizIDByUserName(userName)
	if err != nil {
		utils.HandleGetNumErr(c, "GenQuiz err: "+err.Error())
		return
	}

	if len(quizIDs) == 0 {
		quizIDs = append(quizIDs, 0)
		log.Println("GenQuiz user have no quiz")
	}
	quizID := GetMax(quizIDs) + uint32(1)
	for _, qid := range quizQuestionRst {
		store.GetDB().CommitHistoryRW.Insert(&model.CommitHistory{
			UserName: tt.UserName,
			QuestionID: qid,
			QuizID: quizID,

		})
	}
	utils.HandleGetSuccess(c, model.NewQuizResult{
		QuestionID: quizQuestionRst,
		QuizID:     quizID,
	})
}

func GetMax(vs []uint32) uint32 {
	max := vs[0]
	for _, v := range vs {
		if v > max {
			max = v
		}
	}
	return max
}

// GetQuestionCorrectRateByUser get correct rate in redis
func GetQuestionCorrectRateByUser(userName string, questionID uint32) float32 {
	conn := db.RedisClient.GetConn()
	res, err := conn.Do("HGET", userName, questionID)
	if err != nil {
		log.Println("etQuestionCorrectRateByUser Redis HGET err: ", err)
	}
	f, ok := res.(float32)
	if ok {
		return f
	}
	return float32(1)
}
