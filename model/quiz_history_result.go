package model

type QuizHistoryResult struct {
	QuizID    uint32   `json:"quizId"`
	Types     []uint32 `json:"types"`
	StartTime string   `json:"startTime"`
	CostTime  uint32   `json:"costTime"`
	Point     uint32   `json:"point"`
}
