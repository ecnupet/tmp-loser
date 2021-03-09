package model

import "time"

type CommitHistory struct {
	HistoryID uint64    `xorm:"history_id"`
	UserID    uint64    `xorm:"user_id"`
	QuizID    uint64    `xorm:"quiz_id"`
	Choose    string    `xorm:"choose"`
	TestID    uint64    `xorm:"test_id"`
	Correct   string    `xorm:"correct"`
	CreatedAt time.Time `xorm:"created_at"`
}
