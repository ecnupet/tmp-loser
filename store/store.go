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
	mockStore *MockStore
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

type MockStore struct {
	*mock.MockCommitHistoryReadWriter
	*mock.MockQuestionReadWriter
}

func NewStore() *Store {
	return &Store {
		CommitHistoryRW:  origin.NewCommitHistoryRW(db.Engine),
		QuestionRW: origin.NewQuestionRW(db.Engine),
	}
}

func newMockStore(c *gomock.Controller) *MockStore {
	return &MockStore {
		mock.NewMockCommitHistoryReadWriter(c),
		mock.NewMockQuestionReadWriter(c),
	}
}

func GetDB() *Store {
	return store
}

func GetMockDB() *MockStore {
	return mockStore
}