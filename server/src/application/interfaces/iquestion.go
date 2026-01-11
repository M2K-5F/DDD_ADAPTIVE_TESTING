package interfaces

import (
	"adaptivetesting/src/domain/aggregates/question"
	"adaptivetesting/src/domain/identificators"
	"context"
)

type IQuestionReader interface {
	GetByID(ctx context.Context, id identificators.QuestionID) (*question.Question, error)
}
