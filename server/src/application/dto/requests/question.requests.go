package requests

import (
	"adaptivetesting/src/domain/identificators"
)

type CreateQuestionRequest struct {
	CourseID identificators.CourseID
	Text     string
	TopicID  identificators.TopicID
	Answers  []AnswerRequest
}

type AnswerRequest struct {
	Text      string
	IsCorrect bool
}
