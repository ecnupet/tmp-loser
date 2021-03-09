package model

import "time"

type Quiz struct {
	QuizID      uint64    `xorm:"quiz_id pk autoincr"`
	Description string    `xorm:"description"`
	Type        int       `xorm:"type"`
	Options     string    `xorm:"options"`
	Answer      string    `xorm:"answer"`
	UpdatedAt   time.Time `xorm:"update_at"`
	CreatedAt   time.Time    `xorm:"created_at"`
}

func QuizTableName() string {
	return "quiz"
}