package readwriter

import (
	"ecnu.space/tmp-loser/model"
)

// CommitHistoryReadWriter 声明基于表commit_history提供的功能
type CommitHistoryReadWriter interface {
	// CorrectAndInsert 将提交插入数据库
	Insert(commit *model.CommitHistory) error

	// GetCommitsByUserIDAndQuizID 根据用户名和测验ID提供所有历史提交记录
	GetCommitsByUserNameAndQuizID(userName string, quizID uint64) ([]*model.CommitHistory, error)

	// GetQuizIDByUserNameAndPageNoAndNum 分页查找某个用户的测验ID
	GetQuizIDByUserNameAndPageNoAndNum(userID string, pageNo uint64, num uint64) ([]uint64, error)

	// GetCommitsByQuestionID 获取指定题目ID的所有提交记录
	GetCommitsByQuestionID(quizID uint64) ([]*model.CommitHistory, error)
}
