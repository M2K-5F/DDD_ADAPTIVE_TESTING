package group

import (
	"adaptivetesting/src/domain/aggregates/course"
	"adaptivetesting/src/domain/aggregates/identificators"

	"github.com/google/uuid"
)

func NewGroup(course course.Course, name string, maxStudentCount int) (*Group, error) {
	nameVO, err := NewGroupName(name)
	if err != nil {
		return nil, err
	}
	return &Group{
		name:            nameVO,
		id:              identificators.GroupID(uuid.New()),
		byCourseID:      course.ID(),
		createdByID:     course.CreatedByID(),
		maxStudentCount: maxStudentCount,
		studentCount:    0,
	}, nil
}
