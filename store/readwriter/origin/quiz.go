package origin

import (
	"context"

	"ecnu.space/tmp-loser/model"
	"github.com/go-xorm/xorm"
)

const QuizTableName = "quiz"
// 接口实现
type QuizRW struct {
	engine *xorm.Engine
}

func NewQuizRW(engine *xorm.Engine) *QuizRW {
	return &QuizRW{
		engine: engine,
	}
}

func (rw *QuizRW) TableName() string {
	return QuizTableName
}
// TODO(shanchao)
func (rw *QuizRW) BatchInsert(ctx context.Context, quizs []*model.Quiz) error

// TODO(shanchao)
func (rw *QuizRW) Insert(ctx context.Context, quiz *model.Quiz) error

// TODO(shanchao)
func (rw *QuizRW) GetQuizByType(ctx context.Context, t string) ([]model.Quiz, error)