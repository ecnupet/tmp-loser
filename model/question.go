package model

import "time"

type QuestionType int

const (
	// 传染病
	InfectiousDisease QuestionType = 0
	// 寄生虫病
	ParasiticDisease QuestionType = 1
	// 内科疾病
	InternalDisease QuestionType = 2
	// 外产科疾病
	ObstetricDisease QuestionType = 3
	// 宠物常用手术知识
	SurgicalKnowledge QuestionType = 4
	// 宠物免疫知识
	PetImmunity QuestionType = 5
)

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
