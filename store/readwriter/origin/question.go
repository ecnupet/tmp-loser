package origin

import (
	"context"

	"ecnu.space/tmp-loser/model"
	"github.com/go-xorm/xorm"
)

const QuestionTableName = "question"

// 接口实现
type QuestionRW struct {
	engine *xorm.Engine
}

func NewQuestionRW(engine *xorm.Engine) *QuestionRW {
	return &QuestionRW{
		engine: engine,
	}
}

func (rw *QuestionRW) TableName() string {
	return QuestionTableName
}

func (rw *QuestionRW) BatchInsert(ctx context.Context, questions []*model.Question) error {
	_, err := rw.engine.Table(rw.TableName()).InsertMulti(questions)
	return err
}

func (rw *QuestionRW) Insert(ctx context.Context, question *model.Question) error {
	_, err := rw.engine.Table(rw.TableName()).InsertOne(question)
	return err
}

// TODO(shanchao)
func (rw *QuestionRW) GetQuestionByType(ctx context.Context, t string) ([]model.Question, error)

// TODO(shanchao)
func (rw *QuestionRW) UpdateQuestion(ctx context.Context, question model.Question) error
