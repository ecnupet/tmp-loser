package api

import (
	"fmt"
	"log"
	"sync"

	"ecnu.space/tmp-loser/db"
	"ecnu.space/tmp-loser/model"
	"ecnu.space/tmp-loser/store"
	"ecnu.space/tmp-loser/utils"
	"github.com/gin-gonic/gin"
)

const (
	requiredQuestionNumLowBound = 1
	questionNumPerType          = 5
)

var (
	// map 互斥锁
	mutex           sync.Mutex
	wg              sync.WaitGroup
)

// 鉴权配置前，先假设提供user_name参数
func GenQuiz(c *gin.Context) {
	userNameAny, e := c.Get("user_name")
	if userNameAny == nil && !e {
		utils.HandleGetDBErr(c, "userName quizId 缺一不可")
		return
	}
	userName := userNameAny.(string)
	questionIDSlice := []uint32{}
	log.Println("执行次数：",11)
	tt := model.NewQuizParams{}
	err := c.ShouldBind(&tt)
	if err != nil {
		log.Println("GenQuiz  err:", err)
		utils.HandlePostQuizQuestionErr(c, err.Error())
		return
	}
	ts := tt.Types
	// ratesMap := make(map[uint32]float32)
	fmt.Println("题目类型： ", ts)
	for _, t := range ts {
		log.Println("查询题目类型:", t)
		wg.Add(1)
		go func(t uint32) {
			questions, err := store.GetDB().QuestionRW.GetQuestionByType(t)
			if err != nil {
				log.Println("GenQuiz GetQuestionByType error: ", err)
				utils.HandleGetDBErr(c, err.Error())
				return
			}
			for i := 0; i < questionNumPerType && i < len(questions); i++ {
				mutex.Lock()
				questionIDSlice = append(questionIDSlice, questions[i].QuestionID)
				fmt.Println("questions: ", questionIDSlice)
				mutex.Unlock()
			}
			wg.Done()
		}(t)
	}
	wg.Wait()

	// // 从小到大
	// sort.Float64Slice(rates).Sort()
	if len(questionIDSlice) < requiredQuestionNumLowBound {
		log.Println("GenQuiz shortage in question num")
		utils.HandleGetNumErr(c, "GenQuiz shortage in question num")
		return
	}
	// quizQuestionRst := []uint32{}
	// for k, _ := range ratesMap {
	// 	quizQuestionRst = append(quizQuestionRst, k)
	// }
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
	for _, qid := range questionIDSlice {
		store.GetDB().CommitHistoryRW.Insert(&model.CommitHistory{
			UserName:   tt.UserName,
			QuestionID: qid,
			QuizID:     quizID,
		})
	}
	utils.HandleGetSuccess(c, model.NewQuizResult{
		QuestionID: questionIDSlice,
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

// getMultiRandNum top must over num
// func getMultiRandNum(top int, num int) []int {
// 	if top < num {
// 		return []int{}
// 	}

// }
