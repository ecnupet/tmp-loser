package model

type NewQuizResult struct {
	QuestionID	[]uint32 `json:"questionId"`
	QuizID		uint32 	`json:"quizId"`
}
