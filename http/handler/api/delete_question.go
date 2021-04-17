package api

import (
	"ecnu.space/tmp-loser/model"
	"ecnu.space/tmp-loser/store"
	"ecnu.space/tmp-loser/utils"
	"github.com/gin-gonic/gin"
)

func DeleteQuestion(c *gin.Context) {
	userName := c.GetString("user_name")
	if userName == "" {
		utils.HandleGetErr(c, "userName doesn't exist")
		return
	}
	// isAdmin := c.GetInt("isAdmin")
	// if isAdmin != 1 {
	// 	utils.HandleNotAdminErr(c, "非管理员")
	// 	return
	// }
	question := model.Question{}
	err := c.ShouldBind(&question)
	if err != nil {
		utils.HandlePostQuizQuestionErr(c, "DeleteQuestion bind error: "+err.Error())
		return
	}
	err = store.GetDB().QuestionRW.DeleteQuestion(question.QuestionID)
	if err != nil {
		utils.HandlePostDBErr(c, err.Error())
		return
	}
	utils.HandlePostSuccess(c, "删除成功")
}
