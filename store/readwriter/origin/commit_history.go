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

// Insert insert 暂时不支持事务
func (rw *CommitHistoryRW) Insert(commit *model.CommitHistory) (int32, error) {
	_, err := rw.engine.Table(rw.TableName()).InsertOne(commit)
	if err != nil {
		return -2, err
	}
	maxHistoryID := -1
	_, err = rw.engine.Table(rw.TableName()).Select("max(history_id)").Get(&maxHistoryID)
	if err != nil {
		return -3, err
	}
	return int32(maxHistoryID), nil
}

func (rw *CommitHistoryRW) GetCommitsByUserNameAndQuizID(userName string, quizID uint32) ([]*model.CommitHistory, error) {
	chs := make([]*model.CommitHistory, 0)
	err := rw.engine.Table(rw.TableName()).Where("user_name = ? and quiz_id = ?", userName, quizID).Find(&chs)
	if err != nil {
		return nil, err
	}
	return chs, nil
}

func (rw *CommitHistoryRW) GetQuizIDByUserNameAndPageNoAndNum(userName string, page uint32, pageSize uint32) ([]uint32, error) {
	quizIDs := make([]uint32, 0)
	err := rw.engine.Table(rw.TableName()).Select("quiz_id").Where("user_name = ?", userName).GroupBy("quiz_id").Limit(int(pageSize), int((page-1)*pageSize)).Find(&quizIDs)
	if err != nil {
		return nil, err
	}
	return quizIDs, nil
}
func (rw *CommitHistoryRW) GetAllCommitsHistoryByUserNameAndQuizID(userName string, quizID uint32) ([]*model.CommitHistory, error) {
	chs := make([]*model.CommitHistory, 0)
	err := rw.engine.Table(rw.TableName()).Where("quiz_id = ? and user_name = ?", quizID, userName).Find(&chs)
	if err != nil {
		return nil, err
	}
	return chs, nil
}

func (rw *CommitHistoryRW) GetCommitsByQuestionID(questionID uint32) ([]*model.CommitHistory, error) {
	chs := make([]*model.CommitHistory, 0)
	err := rw.engine.Table(rw.TableName()).Where("question_id = ?", questionID).Find(&chs)
	if err != nil {
		return nil, err
	}
	return chs, err
}

func (rw *CommitHistoryRW) GetCommitByHistoryID(historyID uint32) ([]*model.CommitHistory, error) {
	chs := make([]*model.CommitHistory, 0)
	err := rw.engine.Table(rw.TableName()).Where("history_id = ?", historyID).Find(&chs)
	if err != nil {
		return nil, err
	}
	return chs, nil
}

func (rw *CommitHistoryRW) GetQuizIDByUserName(userName string) ([]uint32, error) {
	qids := []uint32{}
	err := rw.engine.Table(rw.TableName()).Select("quiz_id").Where("user_name = ?", userName).GroupBy("quiz_id").Find(&qids)
	if err != nil {
		return nil, err
	}
	return qids, nil
}

func (rw *CommitHistoryRW) UpdateCommitHistory(q model.CommitHistory) error {
	sql := "update commit_history set choose = ?,spend = ?,correct = ? where user_name = ? and quiz_id = ? and question_id = ?"
	_, err := rw.engine.Table(rw.TableName()).Exec(sql, q.Choose, q.Spend, q.Correct, q.UserName, q.QuizID, q.QuestionID)
	if err != nil {
		return err
	}
	return nil
}
