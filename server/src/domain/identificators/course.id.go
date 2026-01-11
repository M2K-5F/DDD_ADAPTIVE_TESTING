package identificators

import "github.com/google/uuid"

type CourseID uuid.UUID

func (id CourseID) String() string {
	return uuid.UUID(id).String()
}

func NewCourseID() CourseID {
	return CourseID(uuid.New())
}

func ParseCourseID(plain string) (CourseID, error) {
	id, err := uuid.Parse(plain)
	if err != nil {
		return CourseID(uuid.Nil), err
	}

	return CourseID(id), nil
}
