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
	QuestionID  uint32            `json:"questionId" xorm:"PK AUTOINCR 'question_id'"`
	Description string            `json:"description" xorm:"description"`
	Type        uint32            `json:"type" xorm:"type"`
	Options     map[string]string `json:"options" xorm:"options"`
	Answer      string            `json:"answer" xorm:"answer"`
	Duration    uint32            `json:"duration" xorm:"duration"`
	UpdatedAt   time.Time         `json:"updatedAt" xorm:"DateTime UPDATED 'updated_at'"`
	CreatedAt   time.Time         `json:"createdAt" xorm:"DateTime CREATED 'created_at'"`
}
