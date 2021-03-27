package store

import (
	"ecnu.space/tmp-loser/db"
	"ecnu.space/tmp-loser/store/readwriter"
	"ecnu.space/tmp-loser/store/readwriter/origin"
)

var (
	store *Store
)
// 包加载时执行
func init(){
	store = NewStore()
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

func GetDB() *Store {
	return store
}