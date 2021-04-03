package model

type QuizHistoryDetailResultQuestion struct {
	QuestionID  uint32            `json:"questionId"`
	Description string            `json:"description"`
	Type        uint32            `json:"type"`
	Options     map[string]string `json:"options"`
	Duration    uint32            `json:"duration"`
	Answer      string           `json:"answer"`
	Choice      interface{}            `json:"choice"`
	Spend       uint32            `json:"spend"`
}

type QuizHistoryDetailResult struct {
	StartTime string                            `json:"startTime"`
	CostTime  uint32                            `json:"costTime"`
	Results   []QuizHistoryDetailResultQuestion `json:"results"`
}
