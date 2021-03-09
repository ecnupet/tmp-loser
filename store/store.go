package store

import (
	"ecnu.space/tmp-loser/store/readwriter"
	"ecnu.space/tmp-loser/store/readwriter/origin"
	"ecnu.space/tmp-loser/db"
)

type Store struct {
	CommitHistoryRW readwriter.CommitHistoryReadWriter
	QuizRW          readwriter.QuizReadWriter
}

func NewStore() *Store {
	return &Store {
		CommitHistoryRW:  origin.NewCommitHistoryRW(db.Engine),
		QuizRW: origin.NewQuizRW(db.Engine),
	}
}
