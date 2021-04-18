package api

import (
	"strconv"

	"ecnu.space/tmp-loser/model"
	"ecnu.space/tmp-loser/store"
	"ecnu.space/tmp-loser/utils"
	"github.com/gin-gonic/gin"
)

func SearchQuestion(c *gin.Context) {
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
	pageString := c.Query("page")
	pageSizeString := c.Query("pageSize")
	pageSize, err := strconv.Atoi(pageSizeString)
	if err != nil {
		utils.HandleGetErr(c, err.Error())
		return
	}
	page, err := strconv.Atoi(pageString)
	if err != nil {
		utils.HandleGetErr(c, err.Error())
		return
	}
	if keyword == "" {
		count, err := store.GetDB().QuestionRW.GetAllQuestionNum()
		if err != nil {
			utils.HandleGetDBErr(c, err.Error())
			return
		}
		qs, err := store.GetDB().QuestionRW.GetAllQuestionByPage(uint32(page), uint32(pageSize))
		if err != nil {
			utils.HandleGetDBErr(c, err.Error())
			return
		}
		utils.HandleGetSuccess(c, struct {
			Number int64 `json:"count"`
			List   []*model.Question `json:"records"`
		}{
			Number: count,
			List:   qs,
		})
		return
	}
	count, err := store.GetDB().QuestionRW.FuzzySearchNum(keyword)
	if err != nil {
		utils.HandleGetDBErr(c, err.Error())
		return
	}
	qs, err := store.GetDB().QuestionRW.FuzzySearchByPage(keyword, uint32(page), uint32(pageSize))
	if err != nil {
		utils.HandleGetDBErr(c, err.Error())
		return
	}
	utils.HandleGetSuccess(c, struct {
		Number int64 `json:"count"`
		List   []*model.Question 	`json:"records"`
	}{
		Number: count,
		List:   qs,
	})
}
