package identificators

import "github.com/google/uuid"

type EnrollmentID uuid.UUID

func (id EnrollmentID) String() string {
	return uuid.UUID(id).String()
}

func NewEnrollmentID() EnrollmentID {
	return EnrollmentID(uuid.New())
}

func ParseEnrollmentID(plain string) (EnrollmentID, error) {
	id, err := uuid.Parse(plain)
	if err != nil {
		return EnrollmentID(uuid.Nil), err
	}
	return EnrollmentID(id), nil
}
