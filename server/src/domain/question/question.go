package question

import (
	"adaptivetesting/src/domain/question/entities"

	"github.com/google/uuid"
)

type QuestionID uuid.UUID

type Question struct {
	id      QuestionID
	text    string
	answers []entities.Answer
}
