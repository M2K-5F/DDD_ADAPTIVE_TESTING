package group

import (
	"adaptivetesting/src/domain/course"

	"github.com/google/uuid"
)

type GroupID uuid.UUID

type Group struct {
	id              GroupID
	course          course.CourseID
	name            string
	studentCount    int
	maxStudentCount int
}
