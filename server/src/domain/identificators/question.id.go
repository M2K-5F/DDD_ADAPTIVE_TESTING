package identificators

import "github.com/google/uuid"

type QuestionID uuid.UUID

func (id QuestionID) String() string {
	return uuid.UUID(id).String()
}

func NewQuestionID() QuestionID {
	return QuestionID(uuid.New())
}

func ParseQuestionID(plain string) (QuestionID, error) {
	id, err := uuid.Parse(plain)
	if err != nil {
		return QuestionID(uuid.Nil), err
	}
	return QuestionID(id), nil
}
