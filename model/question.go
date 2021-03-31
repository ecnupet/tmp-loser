package model

import "time"

// 映射表 question
type Question struct {
	QuestionID  uint32            `xorm:"PK AUTOINCR 'question_id'"`
	Description string            `xorm:"description"`
	Type        uint32            `xorm:"type"`
	Options     map[string]string `xorm:"options"`
	Answer      string            `xorm:"answer"`
	Duration    uint32            `xorm:"duration"`
	UpdatedAt   time.Time         `xorm:"DateTime UPDATED 'updated_at'"`
	CreatedAt   time.Time         `xorm:"DateTime CREATED 'created_at'"`
}
