package group

import (
	"adaptivetesting/src/domain/aggregates/course"

	"github.com/google/uuid"
)

func NewGroup(course course.Course, name string, maxStudentCount int) (*Group, error) {
	nameVO, err := NewGroupName(name)
	if err != nil {
		return nil, err
	}
	return &Group{
		name:            nameVO,
		id:              GroupID(uuid.New()),
		courseID:        course.ID(),
		maxStudentCount: maxStudentCount,
		studentCount:    0,
	}, nil
}
