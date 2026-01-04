package identificators

import "github.com/google/uuid"

type QuestionID uuid.UUID

func (id QuestionID) String() string {
	return uuid.UUID(id).String()
}
