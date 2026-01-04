package requests

import (
	"adaptivetesting/src/domain/aggregates/identificators"
)

type CreateQuestionRequest struct {
	Text    string
	TopicID identificators.TopicID
	Answers []AnswerRequest
}

type AnswerRequest struct {
	Text      string
	IsCorrect bool
}
