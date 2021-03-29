package model

// 测试试卷的题目，需要返回的题目，接收web端提交的解答json均用这个结构体
type QuizQuestion struct {
	UserName    string            `json:"user_name"`
	QuestionID  uint64            `json:"question_id"`
	Description string            `json:"description"`
	Options     map[string]string `json:"options"`
	Type        uint64            `json:"type"`
	QuizID      uint64            `json:"quiz_id"`
	Order       uint64            `json:"order"`
	// 用户选择 为空
	Choose string `json:"choose"`
	// 花费的时间 为空
	Spend uint64 `json:"spend"`
	// 时间限制
	Duration uint64 `json:"duration"`
}
