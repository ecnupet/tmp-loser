package readwriter

import (
	"ecnu.space/tmp-loser/model"
)

// CommitHistoryReadWriter 声明基于表commit_history提供的功能
type CommitHistoryReadWriter interface {
	// CorrectAndInsert 将提交插入数据库
	Insert(commit *model.CommitHistory) (int32, error)

	// GetCommitsByUserIDAndQuizID 根据用户名和测验ID提供所有历史提交记录
	GetCommitsByUserNameAndQuizID(userName string, quizID uint32) ([]*model.CommitHistory, error)

	// GetQuizIDByUserNameAndPageNoAndNum 分页查找某个用户的测验ID
	GetQuizIDByUserNameAndPageNoAndNum(userName string, pageNo uint32, num uint32) ([]uint32, error)

	// GetAllCommitsHistoryByUserNameAndQuizID 获取用户每次测试的信息
	GetAllCommitsHistoryByUserNameAndQuizID(userName string, quizID uint32) ([]*model.CommitHistory, error)

	// GetQuizIDByUserName 获取用户所有quizID
	GetQuizIDByUserName(userName string)([]uint32, error)

	// GetCommitsByQuestionID 获取指定题目ID的所有提交记录
	GetCommitsByQuestionID(quizID uint32) ([]*model.CommitHistory, error)

	// GetCommitByHistoryID
	GetCommitByHistoryID(historyID uint32) ([]*model.CommitHistory, error)

	// UpdateCommitHistory
	UpdateCommitHistory(qq model.CommitHistory) error
}
