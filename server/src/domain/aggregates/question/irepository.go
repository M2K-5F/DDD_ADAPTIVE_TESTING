package question

import "adaptivetesting/src/domain/aggregates/identificators"

type IRepository interface {
	Save(question *Question) error
	GetByID(id identificators.QuestionID) (*Question, error)
}
