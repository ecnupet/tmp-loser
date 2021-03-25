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

// TODO(shanchao)
func (rw *CommitHistoryRW) CorrectAndInsert(ctx context.Context, commit *model.QuizQuestion) error

// TODO(shanchao)
func (rw *CommitHistoryRW) GetCommitsByUserNameAndQuizID(ctx context.Context, userName string, quizID uint64) []model.CommitHistory

// TODO(shanchao)
func (rw *CommitHistoryRW) GetQuizIDByUserNameAndPageNoAndNum(ctx context.Context, userID string, pageNo uint64, num uint64) []uint64

// TODO(shanchao)
func (rw *CommitHistoryRW) GetCommitsByQuestionID(ctx context.Context, quizID uint64) []model.CommitHistory

// TODO(shanchao)
func (rw *CommitHistoryRW) GenQuiz(userName string) []model.QuizQuestion
