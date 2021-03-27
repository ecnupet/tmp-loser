package api

import (
	"log"
	"sort"

	"ecnu.space/tmp-loser/db"
	"ecnu.space/tmp-loser/store"
	"ecnu.space/tmp-loser/utils"
	"github.com/gin-gonic/gin"
)

const (
	questionNum = 10
)

// 鉴权配置前，先假设提供user_name参数
func GenQuiz(c *gin.Context) {
	userName := c.Query("user_name")
	t := c.Query("type")
	questions, err := store.GetDB().QuestionRW.GetQuestionByType(t)
	if err != nil {
		log.Println("GenQuiz GetQuestionByType error: ", err)
		utils.HandleGetDBErr(c)
		return
	}
	rates := []float64{}
	for _, question := range questions {
		rates = append(rates, GetQuestionCorrectRateByUser(userName, question.QuestionID))
	}
	// 从小到大
	sort.Float64Slice(rates).Sort()
	if len(rates) < questionNum {
		log.Println("shortage in question num")
		utils.HandleGetNumErr(c)
		return
	}
	utils.HandleGetSuccess(c, rates)
	return

}

// GetQuestionCorrectRateByUser get correct rate in redis
func GetQuestionCorrectRateByUser(userName string, questionID uint64) float64 {
	conn := db.RedisClient.GetConn()
	res, err := conn.Do("HGET", userName, questionID)
	if err != nil {
		log.Println("etQuestionCorrectRateByUser Redis HGET err: ", err)
	}
	f, ok := res.(float64)
	if ok {
		return f
	}
	return float64(1)
}
