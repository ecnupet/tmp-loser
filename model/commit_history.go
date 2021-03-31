package model

import "time"

// 映射表 commit_history
type CommitHistory struct {
	HistoryID  uint32 `json:"history_id" xorm:"Int(20) NOT NULL PK AUTOINCR 'history_id'"`
	UserName   string `json:"user_name" xorm:"user_name"`
	QuestionID uint32 `json:"question_id" xorm:"question_id"`
	Choose     string `json:"choose" xorm:"choose"`
	QuizID     uint32 `json:"quiz_id" xorm:"quiz_id"`
	Order      uint32 `json:"order" xorm:"order"`
	// 0: 错误, 1: 正确, 2: 未选择
	Correct   uint32    `json:"correct" xorm:"correct"`
	Spend     uint32    `json:"spend" xorm:"spend"`
	CreatedAt time.Time `json:"created_at" xorm:"DateTime CREATED 'created_at'"`
}
