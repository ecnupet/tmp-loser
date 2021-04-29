package api

import (
	"fmt"
	"strconv"

	"ecnu.space/tmp-loser/store"
	"ecnu.space/tmp-loser/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func GetStatistics(c *gin.Context) {
	userNameAny, e := c.Get("user_name")
	if userNameAny == nil && !e {
		utils.HandleGetDBErr(c, "userName quizId 缺一不可")
		return
	}
	userName := userNameAny.(string)
	chs, err := store.GetDB().CommitHistoryRW.GetUserAllCommitHistory(userName)
	logrus.Info("get statistics relative info: ", chs)
	if err != nil {
		utils.HandleGetDBErr(c, err.Error())
	}
	times := len(chs)
	correctTimes := 0
	faultTimes := 0
	quit := 0
	timeSum := 0
	for _, ch := range chs {
		timeSum += int(ch.Spend)
		if ch.Correct == 1 {
			correctTimes += 1
		} else if ch.Correct == 0 {
			faultTimes += 1
		} else if ch.Correct == 2 {
			quit += 1
		}
	}
	average := float64(timeSum) / float64(times)
	averageRst, err := strconv.ParseFloat(fmt.Sprintf("%.2f", average), 64)
	if err != nil {
		utils.HandleGetErr(c, err.Error())
	}

	utils.HandleGetSuccess(c, struct {
		TotalCommitCount   uint32  `json:"totalCommitCount"`
		TotalCorrectCount  uint32  `json:"totalCorrectCount"`
		TotalWrongCount    uint32  `json:"totalWrongCount"`
		TotalNoAnswerCount uint32  `json:"totalNoAnswerCount"`
		AverageAnswerTime  float64 `json:"averageAnswerTime"`
	}{
		TotalCommitCount:   uint32(times),
		TotalCorrectCount:  uint32(correctTimes),
		TotalWrongCount:    uint32(faultTimes),
		TotalNoAnswerCount: uint32(quit),
		AverageAnswerTime:  averageRst,
	})
}
