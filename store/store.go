package store

import (
	"ecnu.space/tmp-loser/db"
	"ecnu.space/tmp-loser/store/readwriter"
	"ecnu.space/tmp-loser/store/readwriter/mock"
	"ecnu.space/tmp-loser/store/readwriter/origin"
	"github.com/golang/mock/gomock"
)

var (
	store *Store
	mockStore *Store
)
// 包加载时执行
func init(){
	store = NewStore()
}

func InitMockClient(c *gomock.Controller) {
	mockStore = newMockStore(c)
}
// 持久层mysql服务
type Store struct {
	CommitHistoryRW readwriter.CommitHistoryReadWriter
	QuestionRW          readwriter.QuestionReadWriter
}

func NewStore() *Store {
	return &Store {
		CommitHistoryRW:  origin.NewCommitHistoryRW(db.Engine),
		QuestionRW: origin.NewQuestionRW(db.Engine),
	}
}

func newMockStore(c *gomock.Controller) *Store {
	return &Store {
		CommitHistoryRW: mock.NewMockCommitHistoryReadWriter(c),
		QuestionRW: mock.NewMockQuestionReadWriter(c),
	}
}

func GetDB() *Store {
	if mockStore != nil {
		return mockStore
	}
	return store
}