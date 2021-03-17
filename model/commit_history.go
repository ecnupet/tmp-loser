package model

import "time"

type CommitHistory struct {
	HistoryID uint64    `xorm:"history_id"`
	UserID    uint64    `xorm:"user_id"`
	QuestionID    uint64    `xorm:"question_id"`
	Choose    string    `xorm:"choose"`
	QuizID    uint64    `xorm:"quiz_id"`
	Correct   string    `xorm:"correct"`
	CreatedAt time.Time `xorm:"created_at"`
}
