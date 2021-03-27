package readwriter

import (
	"ecnu.space/tmp-loser/model"
)

// QuestionReadWriter 声明基于表question提供的功能
type QuestionReadWriter interface {
	// BatchInsert 批量插入题目
	BatchInsert(quizs []*model.Question) error

	// Insert 单个插入题目
	Insert(question *model.Question) error

	// GetQuestionByType根据类型获取所有的题目
	GetQuestionByType(t string) ([]*model.Question, error)

	// UpdateQuestion 更新题目
	UpdateQuestion(questionID uint64, question *model.Question) error
}
