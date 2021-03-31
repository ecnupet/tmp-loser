package model

// 映射表 question
type QuestionStringTime struct {
	QuestionID  uint32            `json:""questionId`
	Description string            `json:"description"`
	Type        uint32            `json:"type"`
	Options     map[string]string `json:"options"`
	Duration    uint32            `json:"duration"`
}
