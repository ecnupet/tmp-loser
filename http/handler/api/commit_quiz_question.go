package api

import (
	"ecnu.space/tmp-loser/model"
	"ecnu.space/tmp-loser/store"
	"ecnu.space/tmp-loser/utils"
	"github.com/gin-gonic/gin"
)

// CommitQuizQUestion 批改提交的指定题目的回答数据，并插入到数据库
func CommitQuizQuestion(c *gin.Context){
	qq := model.QuizQuestion{}
	err := c.ShouldBind(&qq)
	if err != nil {
		utils.HandlePostQuizQuestionErr(c, err.Error())
		return
	}
	rst, err := correct(qq)
	if err != nil {
		utils.HandlePostDBErr(c, err.Error())
		return
	}
	ch := model.CommitHistory{
		QuestionID: qq.QuestionID,
		// 临时设置，之后用person-manage grpc服务处理用户的token, 拿到用户名
		UserName: "sc",
		Choose: qq.Choose,
		QuizID: qq.QuizID,
		Order: qq.Order,
		Spend: qq.Spend,
		Correct: rst,
	}
	chID, err := store.GetDB().CommitHistoryRW.Insert(&ch)
	if err != nil {
		utils.HandlePostDBErr(c, err.Error())
		return
	}
	ch2, err := store.GetDB().CommitHistoryRW.GetCommitByHistoryID(uint64(chID))
	if err != nil {
		utils.HandleGetDBErr(c, err.Error())
	}
	utils.HandlePostSuccess(c, *ch2)
}

// correct 批改提交的回答
func correct(qq model.QuizQuestion) (uint64, error){
	questionID := qq.QuestionID
	question, err := store.GetDB().QuestionRW.GetQuestionById(questionID)
	if err != nil {
		return 2, err
	}
	if qq.Choose == "" {
		return 2, nil
	}else if qq.Choose == question.Answer {
		return 1, nil
	}
	return 0, nil

}