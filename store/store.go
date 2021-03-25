package store

import (
	"ecnu.space/tmp-loser/db"
	"ecnu.space/tmp-loser/store/readwriter"
	"ecnu.space/tmp-loser/store/readwriter/origin"
)

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
