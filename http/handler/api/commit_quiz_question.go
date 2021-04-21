package api

import (
	"errors"
	"log"

	"ecnu.space/tmp-loser/model"
	"ecnu.space/tmp-loser/store"
	"ecnu.space/tmp-loser/utils"
	"github.com/gin-gonic/gin"
)

// CommitQuizQUestion 批改提交的指定题目的回答数据，并插入到数据库
func CommitQuizQuestion(c *gin.Context) {
	userNameAny, e := c.Get("user_name")
	if userNameAny == nil && !e {
		utils.HandleGetDBErr(c, "userName quizId 缺一不可")
		return

	}
	userName := userNameAny.(string)

	cqp := model.CheckQuestinoParams{}
	err := c.ShouldBind(&cqp)
	if err != nil {
		log.Println("CommitQuizQuestion 1 err:", err)
		utils.HandlePostQuizQuestionErr(c, err.Error())
		return
	}
	cqp.UserName = userName
	if cqp.UserName == "" || cqp.QuestinoID == 0 || cqp.QuizID == 0 {
		utils.HandlePostQuizQuestionErr(c, "userName questionId quizId 缺一不可")
		return
	}

	rst, err := correct(cqp)
	if err != nil {
		log.Println("CommitQuizQuestion 2 err:", err)
		utils.HandlePostDBErr(c, err.Error())
		return
	}
	questions, _ := store.GetDB().QuestionRW.GetQuestionById(cqp.QuestinoID)
	if cqp.TimeSpend > questions[0].Duration {
		cqp.TimeSpend = 0
	}
	ch := model.CommitHistory{
		QuestionID: cqp.QuestinoID,
		// 临时设置，之后用person-manage grpc服务处理用户的token, 拿到用户名
		UserName: cqp.UserName,
		Choose:   cqp.Answer,
		QuizID:   cqp.QuizID,
		Spend:    cqp.TimeSpend,
		Correct:  rst,
	}
	log.Println("cqp timeSpend: ", cqp.TimeSpend)

	err = store.GetDB().CommitHistoryRW.UpdateCommitHistory(ch)
	if err != nil {
		log.Println("CommitQuizQuestion 3 err:", err)
		utils.HandlePostDBErr(c, err.Error())
		return
	}
	cor := false
	if rst == 1 {
		cor = true
	}
	utils.HandlePostSuccess(c, model.CheckQuestionResult{
		Correct: cor,
	})
}

// correct 批改提交的回答
func correct(qq model.CheckQuestinoParams) (uint32, error) {
	questionID := qq.QuestinoID
	questions, err := store.GetDB().QuestionRW.GetQuestionById(questionID)
	if err != nil {
		return 2, err
	}
	if len(questions) == 0 {
		return 2, errors.New("该问题不存在")
	}
	if qq.Answer == "" {
		return 2, nil
	} else if qq.Answer == questions[0].Answer {
		return 1, nil
	}
	return 0, nil
}
