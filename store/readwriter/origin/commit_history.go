package origin

import (
	"context"

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

func (rw *CommitHistoryRW) Insert(ctx context.Context, commit *model.CommitHistory) error {
	_, err := rw.engine.Table(rw.TableName()).InsertOne(commit)
	if err != nil {
		return err
	}
	return nil
}

func (rw *CommitHistoryRW) GetCommitsByUserNameAndQuizID(ctx context.Context, userName string, quizID uint64) ([]*model.CommitHistory, error) {
	cs := make([]*model.CommitHistory, 0)
	err := rw.engine.Table(rw.TableName()).Where("user_name = ? and quiz_id = ?", userName, quizID).Find(&cs)
	if err != nil {
		return nil, err
	}
	return cs, nil
}

func (rw *CommitHistoryRW) GetQuizIDByUserNameAndPageNoAndNum(ctx context.Context, userName string, page uint64, pageSize uint64) ([]uint64, error) {
	quizIDs := make([]uint64, 0)
	err := rw.engine.Table(rw.TableName()).Select("quiz_id").Where("user_name = ?", userName).Limit(int(pageSize), int((page-1)*pageSize)).Find(&quizIDs)
	if err != nil {
		return nil, err
	}
	return quizIDs, nil
}

func (rw *CommitHistoryRW) GetCommitsByQuestionID(ctx context.Context, questionID uint64) ([]*model.CommitHistory, error) {
	cs := make([]*model.CommitHistory, 0)
	err := rw.engine.Table(rw.TableName()).Where("question_id = ?", questionID).Find(&cs)
	if err != nil {
		return nil, err
	}
	return cs, err
}
