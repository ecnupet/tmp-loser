mock_gen:
	mockgen  -source=store/readwriter/commit_history.go -destination=store/readwriter/mock/commit_history_mock.go -package=mock
	mockgen  -source=store/readwriter/question.go -destination=store/readwriter/mock/question_mock.go -package=mock