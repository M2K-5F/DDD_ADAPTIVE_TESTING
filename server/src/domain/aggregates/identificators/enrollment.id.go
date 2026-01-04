package identificators

import "github.com/google/uuid"

type EnrollmentID uuid.UUID

func (id EnrollmentID) String() string {
	return uuid.UUID(id).String()
}
