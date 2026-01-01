package topic

import (
	"adaptivetesting/src/domain/course"

	"github.com/google/uuid"
)

type TopicID uuid.UUID

type Topic struct {
	id         TopicID
	courseID   course.CourseID
	Name       string
	isArchived bool

	totalQuestions int
}
