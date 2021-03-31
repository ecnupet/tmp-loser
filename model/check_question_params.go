package model

type CheckQuestinoParams struct {
	UserName   string `json:"userName"`
	QuizID     uint32 `json:"quizId"`
	QuestinoID uint32 `json:"questionId"`
	Answer     string `json:"answer"`
	TimeSpend  uint32 `json:"timeSpend"`
}
