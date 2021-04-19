package api

import (
	"fmt"
	"testing"
	"time"

	"ecnu.space/tmp-loser/model"
	"ecnu.space/tmp-loser/store"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetQuestionDetail(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()
	store.InitMockClient(c)
	mockStore := store.GetMockDB()

	mockStore.MockQuestionReadWriter.EXPECT().GetQuestionById(gomock.Any()).Return([]*model.Question{
		
	})

	mockStore.MockQuestionReadWriter.EXPECT().GetQuestionById(gomock.Any()).Return([]*model.Question{
		{
			QuestionID: 5,
			Description: "慢性肝炎的原因不包括",
			Type: 0,
			Options: map[string]string{
				"A": "甲型肝炎",
				"B": "乙型肝炎",
				"C": "丙型肝炎",
				"D": "丁型肝炎",
			},
			Answer: "A",
			Duration: uint32(10),
			UpdatedAt: time.Now(),
			CreatedAt: time.Now(),
		},
	}, nil)
	ttt := true
	rst, err := mockStore.GetQuestionById(4)
	assert.NoError(t, err)
	assert.True(t, ttt)
	fmt.Println(rst[0])
}