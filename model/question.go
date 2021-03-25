package model

import "time"

// 映射表 question
type Question struct {
	QuestionID  uint64            `xorm:"PK AUTOINCR 'quesion_id'"`
	Description string            `xorm:"description"`
	Type        uint64            `xorm:"type"`
	Options     map[string]string `xorm:"options"`
	Answer      string            `xorm:"answer"`
	Duration    uint64            `xorm:"duration"`
	UpdatedAt   time.Time         `xorm:"DateTime UPDATED 'update_at'"`
	CreatedAt   time.Time         `xorm:"DateTime CREATED 'created_at'"`
}
