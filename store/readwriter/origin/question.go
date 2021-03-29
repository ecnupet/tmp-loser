package origin

import (
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

func (rw *QuestionRW) BatchInsert(questions []*model.Question) error {
	_, err := rw.engine.Table(rw.TableName()).InsertMulti(questions)
	return err
}

func (rw *QuestionRW) Insert(question *model.Question) error {
	_, err := rw.engine.Table(rw.TableName()).InsertOne(question)
	return err
}

func (rw *QuestionRW) GetQuestionByType(t string) ([]*model.Question, error) {
	questions := make([]*model.Question, 0)
	err := rw.engine.Table(rw.TableName()).Where("type = ?", t).Find(&questions)
	if err != nil {
		return nil, err
	}
	return questions, nil
}

func (rw *QuestionRW) UpdateQuestion(questionID uint64, question *model.Question) error {
	_, err := rw.engine.Table(rw.TableName()).ID(questionID).AllCols().Update(question)
	return err
}

func (rw *QuestionRW) GetQuestionById(questionID uint64) ([]*model.Question, error) {
	qs := make([]*model.Question, 0)
	err := rw.engine.Table(rw.TableName()).Where("question_id = ?", questionID).Find(&qs)
	if err != nil {
		return nil, err
	}
	return qs, nil
}
