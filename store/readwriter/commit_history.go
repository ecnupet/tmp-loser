package readwriter

import "ecnu.space/tmp-loser/model"

// CommitHistoryReadWriter 声明基于表commit_history提供的功能
type CommitHistoryReadWriter interface {
	// CorrectAndInsert 批改用户提交的回答并插入数据库
	CorrectAndInsert(commit *model.CommitHistory) error
	// GetCommitsByUserIDAndTestID 根据用户ID个测验ID提供所有历史提交记录
	GetCommitsByUserIDAndTestID() model.CommitHistory
	// GetCommitsByQuizID 获取指定题目ID的所有提交记录
	GetCommitsByQuizID(quizID uint64) []model.Quiz
}