package course

import (
	"adaptivetesting/src/domain/user"

	"github.com/google/uuid"
)

type CourseID uuid.UUID

type Course struct {
	id          CourseID
	createdByID user.UserID
	name        string
	isArchived  bool

	topicCount int
	groupCount int
}

func (c *Course) ID() CourseID {
	return c.id
}
