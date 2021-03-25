package readwriter

import (
	"context"

	"ecnu.space/tmp-loser/model"
)

// QuestionReadWriter 声明基于表question提供的功能
type QuestionReadWriter interface {
	// BatchInsert 批量插入题目
	BatchInsert(ctx context.Context, quizs []*model.Question) error
	// Insert 单个插入题目
	Insert(ctx context.Context, question *model.Question) error
	// GetQuestionByType根据类型获取所有的题目
	GetQuestionByType(ctx context.Context, t string) ([]model.Question, error)
}
