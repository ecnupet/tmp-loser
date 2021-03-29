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

// Insert 插入并返回主键，支持事务，可高并发
func (rw *CommitHistoryRW) Insert(commit *model.CommitHistory) (int64, error) {
	session := rw.engine.Table(rw.TableName())
	// 事务启动
	if err := session.Begin(); err != nil {
		return -1, err
	}
	_, err := session.InsertOne(commit)
	if err != nil {
		return -1, err
	}
	maxHistoryID := -1
	_, err = rw.engine.Table(rw.TableName()).Select("max(history_id)").Get(&maxHistoryID)
	if err != nil {
		return -1, err
	}
	// 事务提交
	if err := session.Commit(); err != nil {
		return -1, err
	}
	return int64(maxHistoryID), nil
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

func (rw *CommitHistoryRW) GetCommitByHistoryID(historyID uint64) ([]*model.CommitHistory, error) {
	chs := make([]*model.CommitHistory, 0)
	err := rw.engine.Table(rw.TableName()).Where("history_id = ?", historyID).Find(&chs)
	if err != nil {
		return nil, err
	}
	return chs, nil
}
