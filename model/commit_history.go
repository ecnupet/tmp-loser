package model

import "time"

// 映射表 commit_history
type CommitHistory struct {
	HistoryID  uint64    `xorm:"Int(20) NOT NULL PK AUTOINCR 'history_id'"`
	UserName   string    `xorm:"user_name"`
	QuestionID uint64    `xorm:"question_id"`
	Choose     string    `xorm:"choose"`
	QuizID     uint64    `xorm:"quiz_id"`
	Order      uint64    `xorm:"order"`
	// 0: 错误, 1: 正确, 2: 未选择
	Correct    uint64    `xorm:"correct"`
	Spend      uint64    `xorm:"spend"`
	CreatedAt  time.Time `xorm:"DateTime CREATED 'created_at'"`
}
