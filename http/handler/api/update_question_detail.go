package api

import (
	"ecnu.space/tmp-loser/model"
	"ecnu.space/tmp-loser/store"
	"ecnu.space/tmp-loser/utils"
	"github.com/gin-gonic/gin"
)

func UpdateQuestionDetail(c *gin.Context) {
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
	q := model.Question{}
	err := c.ShouldBind(&q)
	if err != nil {
		utils.HandlePostDBErr(c, err.Error())
		return 
	}
	err = store.GetDB().QuestionRW.UpdateQuestion(q.QuestionID, &q)
	if err != nil {
		utils.HandlePostDBErr(c, err.Error())
		return 
	}
	utils.HandlePostSuccess(c, "更新成功")
}
