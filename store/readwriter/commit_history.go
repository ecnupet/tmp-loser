package readwriter

type AnswerHistoryReadWriter interface {
	CorrectAndInsert(quizID uint64, userID uint64, )
}