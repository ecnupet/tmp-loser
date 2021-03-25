package readwriter

import (
	"context"

	"ecnu.space/tmp-loser/model"
)

// CommitHistoryReadWriter 声明基于表commit_history提供的功能
type CommitHistoryReadWriter interface {
	// CorrectAndInsert 批改用户提交的回答并插入数据库
	CorrectAndInsert(ctx context.Context, commit *model.QuizQuestion) error

	// GetCommitsByUserIDAndQuizID 根据用户名和测验ID提供所有历史提交记录
	GetCommitsByUserNameAndQuizID(ctx context.Context, userName string, quizID uint64) []model.CommitHistory

	// GetQuizIDByUserNameAndPageNoAndNum 分页查找某个用户的测验ID
	GetQuizIDByUserNameAndPageNoAndNum(ctx context.Context, userID string, pageNo uint64, num uint64) []uint64

	// GetCommitsByQuestionID 获取指定题目ID的所有提交记录
	GetCommitsByQuestionID(ctx context.Context, quizID uint64) []model.CommitHistory

	// TODO(shanchao): 根据用户该题目以往正确率或该题总正确率生成试卷
	// GenQuiz 生成测试试卷
	GenQuiz(userName string) []model.QuizQuestion
}