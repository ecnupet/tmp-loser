package readwriter

import (
	"context"

	"ecnu.space/tmp-loser/model"
)

// QuizReadWriter 声明基于表quiz提供的功能
type QuizReadWriter interface {
	// BatchInsert 批量插入题目
	BatchInsert(ctx context.Context, quizs []*model.Quiz) error
	// Insert 单个插入题目
	Insert(ctx context.Context, quiz *model.Quiz) error
	// GetQuizByType根据类型获取所有的题目
	GetQuizByType(ctx context.Context, t string) ([]model.Quiz, error)
}
