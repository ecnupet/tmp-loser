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

func (rw *QuestionRW) GetQuestionByType(t uint32) ([]*model.Question, error) {
	questions := make([]*model.Question, 0)
	err := rw.engine.Table(rw.TableName()).Where("type = ?", t).Find(&questions)
	if err != nil {
		return nil, err
	}
	return questions, nil
}

func (rw *QuestionRW) UpdateQuestion(questionID uint32, question *model.Question) error {
	_, err := rw.engine.Table(rw.TableName()).ID(questionID).AllCols().Update(question)
	return err
}

func (rw *QuestionRW) GetQuestionById(questionID uint32) ([]*model.Question, error) {
	qs := make([]*model.Question, 0)
	err := rw.engine.Table(rw.TableName()).Where("question_id = ?", questionID).Find(&qs)
	if err != nil {
		return nil, err
	}
	return qs, nil
}

func (rw *QuestionRW) FuzzySearchByPage(keyword string, page, pageSize uint32) ([]*model.Question, error) {
	qs := make([]*model.Question, 0)
	err := rw.engine.Table(rw.TableName()).Where("description LIKE '%"+keyword+"%'").Limit(int(pageSize), int((page-1)*pageSize)).Find(&qs)
	if err != nil {
		return nil, err
	}
	return qs, nil
}

func (rw *QuestionRW) FuzzySearchNum(keyword string) (count int64, err error) {
	count, err = rw.engine.Table(rw.TableName()).Where("description LIKE '%"+keyword+"%'").Count()
	return
}

func (rw *QuestionRW) GetAllQuestionNum() (count int64, err error) {
	count, err = rw.engine.Table(rw.TableName()).Count()
	return
}

func (rw *QuestionRW) GetAllQuestionByPage(page, pageSize uint32) ([]*model.Question, error) {
	qs := make([]*model.Question, 0)
	err := rw.engine.Table(rw.TableName()).Limit(int(pageSize), int((page-1)*pageSize)).Find(qs)
	if err !=nil {
		return nil ,err
	}
	return qs, nil
}

func (rw *QuestionRW) DeleteQuestion(questionId uint32)error {
	q := model.Question{
		QuestionID: questionId,
	}
	_, err := rw.engine.Table(rw.TableName()).Delete(&q)
	if err != nil {
		return err
	}
	return nil
}