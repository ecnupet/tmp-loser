package model

import "time"

// 映射表 commit_history
type CommitHistory struct {
	HistoryID  uint64 `json:"history_id" xorm:"Int(20) NOT NULL PK AUTOINCR 'history_id'"`
	UserName   string `json:"user_name" xorm:"user_name"`
	QuestionID uint64 `json:"question_id" xorm:"question_id"`
	Choose     string `json:"choose" xorm:"choose"`
	QuizID     uint64 `json:"quiz_id" xorm:"quiz_id"`
	Order      uint64 `json:"order" xorm:"order"`
	// 0: 错误, 1: 正确, 2: 未选择
	Correct   uint64    `json:"correct" xorm:"correct"`
	Spend     uint64    `json:"speed" xorm:"spend"`
	CreatedAt time.Time `json:"created_at" xorm:"DateTime CREATED 'created_at'"`
}
