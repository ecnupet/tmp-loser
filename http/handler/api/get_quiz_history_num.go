package api

import (
	"log"
	"ecnu.space/tmp-loser/store"
	"ecnu.space/tmp-loser/utils"
	"github.com/gin-gonic/gin"
)

func GetQuizHistoryNum(c *gin.Context) {
	userNameAny, _ := c.Get("user_name")
	userName := userNameAny.(string)
	if userName == "" {
		utils.HandleGetErr(c, "userName doesn't exist")
		return
	}
	ids, err := store.GetDB().CommitHistoryRW.GetQuizIDByUserName(userName)
	log.Println(ids)
	if err != nil {
		utils.HandleGetDBErr(c, "GetQuizHistoryNum err: "+err.Error())
		return
	}
	utils.HandleGetSuccess(c, struct {
		Number uint32 `json:"number"`
	}{
		Number: uint32(len(ids)),
	})
}
