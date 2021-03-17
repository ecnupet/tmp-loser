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
func (rw *CommitHistoryRW) CorrectAndInsert(ctx context.Context, commit *model.CommitHistory) error

// TODO(shanchao)
func (rw *CommitHistoryRW) GetCommitsByUserIDAndTestID(ctx context.Context) model.CommitHistory

// TODO(shanchao)
func (rw *CommitHistoryRW) GetCommitsByQuestionID(ctx context.Context, questionID uint64) []model.Question
