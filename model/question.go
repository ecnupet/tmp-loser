package model

import "time"

type Question struct {
	QuestionID      uint64    `xorm:"quesion_id pk autoincr"`
	Description string    `xorm:"description"`
	Type        int       `xorm:"type"`
	Options     string    `xorm:"options"`
	Answer      string    `xorm:"answer"`
	Duration	time.Time	`xorm:"duration"`
	UpdatedAt   time.Time `xorm:"update_at"`
	CreatedAt   time.Time    `xorm:"created_at"`
}
