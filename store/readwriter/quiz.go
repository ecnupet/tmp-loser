package readwriter

import "ecnu.space/tmp-loser/model"

// QuizReadWriter 声明基于表quiz提供的功能
type QuizReadWriter interface {
	// BatchInsert 批量插入题目
	BatchInsert(quizs []*model.Quiz) error
	// Insert 单个插入题目
	Insert(quiz *model.Quiz) error
	// GetQuizByType根据类型获取所有的题目
	GetQuizByType(t string) ([]model.Quiz, error)

}