package api

import (
	"ecnu.space/tmp-loser/store"
	"ecnu.space/tmp-loser/utils"
	"github.com/gin-gonic/gin"
)

func SearchQuestionNum(c *gin.Context) {
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
	keyword := c.Query("keyword")
	if keyword == "" {
		count, err := store.GetDB().QuestionRW.GetAllQuestionNum()
		if err != nil {
			utils.HandleGetDBErr(c, err.Error())
			return
		}
		utils.HandleGetSuccess(c, struct {
			Number uint32 `json:"number"`
		}{
			Number: uint32(count),
		})
		return
	}
	count, err := store.GetDB().QuestionRW.FuzzySearchNum(keyword)
	if err != nil {
		utils.HandleGetDBErr(c, err.Error())
		return
	}
	utils.HandleGetSuccess(c, struct {
		Number uint32 `json:"number"`
	}{
		Number: uint32(count),
	})
}
