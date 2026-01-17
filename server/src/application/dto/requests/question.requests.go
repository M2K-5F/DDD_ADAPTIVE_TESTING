package requests

import (
	"adaptivetesting/src/domain/aggregates/question"
	"adaptivetesting/src/domain/identificators"
)

type CreateQuestionRequest struct {
	CourseID identificators.CourseID
	Text     string
	TopicID  identificators.TopicID
	Answers  []*question.AnswerCreate
}
