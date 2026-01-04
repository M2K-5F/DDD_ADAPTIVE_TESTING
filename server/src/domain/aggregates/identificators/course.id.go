package identificators

import "github.com/google/uuid"

type CourseID uuid.UUID

func (id CourseID) String() string {
	return uuid.UUID(id).String()
}
