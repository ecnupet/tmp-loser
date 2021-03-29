package origin

import (
	"ecnu.space/tmp-loser/model"
	"github.com/go-xorm/xorm"
)

const CommitHistoryTableName = "commit_history"

// 接口实现
type CommitHistoryRW struct {
	engine *xorm.Engine
}

func NewCommitHistoryRW(engine *xorm.Engine) *CommitHistoryRW {
	return &CommitHistoryRW{
		engine: engine,
	}
}

func (rw *CommitHistoryRW) TableName() string {
	return CommitHistoryTableName
}

func (rw *CommitHistoryRW) Insert(commit *model.CommitHistory) (int64, error) {
	// 返回主键
	p, err := rw.engine.Table(rw.TableName()).InsertOne(commit)
	if err != nil {
		return p, err
	}
	return p, nil
}

func (rw *CommitHistoryRW) GetCommitsByUserNameAndQuizID(userName string, quizID uint64) ([]*model.CommitHistory, error) {
	chs := make([]*model.CommitHistory, 0)
	err := rw.engine.Table(rw.TableName()).Where("user_name = ? and quiz_id = ?", userName, quizID).Find(&chs)
	if err != nil {
		return nil, err
	}
	return chs, nil
}

func (rw *CommitHistoryRW) GetQuizIDByUserNameAndPageNoAndNum(userName string, page uint64, pageSize uint64) ([]uint64, error) {
	quizIDs := make([]uint64, 0)
	err := rw.engine.Table(rw.TableName()).Select("quiz_id").Where("user_name = ?", userName).Limit(int(pageSize), int((page-1)*pageSize)).Find(&quizIDs)
	if err != nil {
		return nil, err
	}
	return quizIDs, nil
}

func (rw *CommitHistoryRW) GetCommitsByQuestionID(questionID uint64) ([]*model.CommitHistory, error) {
	chs := make([]*model.CommitHistory, 0)
	err := rw.engine.Table(rw.TableName()).Where("question_id = ?", questionID).Find(&chs)
	if err != nil {
		return nil, err
	}
	return chs, err
}

func (rw * CommitHistoryRW) GetCommitByHistoryID(historyID uint64) (*model.CommitHistory, error) {
	ch := new(model.CommitHistory)
	err := rw.engine.Table(rw.TableName()).Where("history_id = ?", historyID).Find(&ch)
	if err != nil {
		return nil ,err
	}
	return ch, nil
}