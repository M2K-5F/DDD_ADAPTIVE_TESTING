package group

import (
	"adaptivetesting/src/domain/aggregates/course"

	"github.com/google/uuid"
)

type GroupID uuid.UUID

type Group struct {
	id              GroupID
	courseID        course.CourseID
	name            GroupName
	studentCount    int
	maxStudentCount int
}

func (g *Group) ID() GroupID {
	return g.id
}
