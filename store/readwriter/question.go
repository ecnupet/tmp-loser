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
	GetQuestionByType(t uint32) ([]*model.Question, error)

	// UpdateQuestion 更新题目
	UpdateQuestion(questionID uint32, question *model.Question) error

	// GetQuestionById get question by id
	GetQuestionById(questionID uint32) ([]*model.Question, error)

	// FuzzySearch 根据关键字模糊查询
	FuzzySearchByPage(keyword string, page, pageSize uint32)([]*model.Question, error) 

	// FuzzySearchNum ...
	FuzzySearchNum(keyword string)(int64, error)

	// GetAllQuestion
	GetAllQuestionNum()(int64, error)

	// GetAllQuestionByPage
	GetAllQuestionByPage(page, pageSize uint32)([]*model.Question, error)

	// DeleteQuestionById
	DeleteQuestion(questionId uint32)error
}
