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
	qq := model.QuizQuestion{}
	err := c.ShouldBind(&qq)
	log.Println(qq)
	if err != nil {
		log.Println("CommitQuizQuestion 1 err:", err)
		utils.HandlePostQuizQuestionErr(c, err.Error())
		return
	}
	rst, err := correct(qq)
	if err != nil {
		log.Println("CommitQuizQuestion 2 err:", err)
		utils.HandlePostDBErr(c, err.Error())
		return
	}
	ch := model.CommitHistory{
		QuestionID: qq.QuestionID,
		// 临时设置，之后用person-manage grpc服务处理用户的token, 拿到用户名
		UserName: "sc",
		Choose:   qq.Choose,
		QuizID:   qq.QuizID,
		Order:    qq.Order,
		Spend:    qq.Spend,
		Correct:  rst,
	}
	chID, err := store.GetDB().CommitHistoryRW.Insert(&ch)
	if err != nil {
		log.Println("CommitQuizQuestion 3 err:", err)
		utils.HandlePostDBErr(c, err.Error())
		return
	}
	chs2, err := store.GetDB().CommitHistoryRW.GetCommitByHistoryID(uint64(chID))
	if err != nil {
		log.Println("CommitQuizQuestion 4 err:", err)
		utils.HandleGetDBErr(c, err.Error())
	}
	if len(chs2) < 1 {
		log.Println("CommitQuizQuestion 4 err:", err)
		utils.HandlePostQuizQuestionErr(c, "无commitHistory")
		return
	}
	utils.HandlePostSuccess(c, *chs2[0])
}

// correct 批改提交的回答
func correct(qq model.QuizQuestion) (uint64, error) {
	questionID := qq.QuestionID
	questions, err := store.GetDB().QuestionRW.GetQuestionById(questionID)
	if err != nil {
		return 2, err
	}
	if len(questions) == 0 {
		return 2, errors.New("该问题不存在")
	}
	if qq.Choose == "" {
		return 2, nil
	} else if qq.Choose == questions[0].Answer {
		return 1, nil
	}
	return 0, nil

}
