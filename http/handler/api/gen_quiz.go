package api

import (
	"log"

	"ecnu.space/tmp-loser/db"
	"ecnu.space/tmp-loser/store"
	"ecnu.space/tmp-loser/utils"
	"github.com/gin-gonic/gin"
)

const (
	requiredQuestionNum = 1
)

// 鉴权配置前，先假设提供user_name参数
func GenQuiz(c *gin.Context) {
	userName := c.Query("user_name")
	t := c.Query("type")
	questions, err := store.GetDB().QuestionRW.GetQuestionByType(t)
	if err != nil {
		log.Println("GenQuiz GetQuestionByType error: ", err)
		utils.HandleGetDBErr(c, err.Error())
		return
	}
	ratesMap := make(map[uint32]float32)
	for _, question := range questions {
		ratesMap[question.QuestionID] = GetQuestionCorrectRateByUser(userName, question.QuestionID)
	}
	// // 从小到大
	// sort.Float64Slice(rates).Sort()
	if len(ratesMap) < requiredQuestionNum {
		log.Println("GenQuiz shortage in question num")
		utils.HandleGetNumErr(c, "GenQuiz shortage in question num")
		return
	}
	utils.HandleGetSuccess(c, ratesMap)
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
