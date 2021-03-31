package model

// 测试试卷的题目，需要返回的题目，接收web端提交的解答json均用这个结构体
type QuizQuestion struct {
	UserName    string            `json:"user_name"`
	QuestionID  uint32            `json:"question_id"`
	Description string            `json:"description"`
	Options     map[string]string `json:"options"`
	Type        uint32            `json:"type"`
	QuizID      uint32            `json:"quiz_id"`
	Order       uint32            `json:"order"`
	// 用户选择 为空
	Choose string `json:"choose"`
	// 花费的时间 为空
	Spend uint32 `json:"spend"`
	// 时间限制
	Duration uint32 `json:"duration"`
}
