package course

import (
	"adaptivetesting/src/domain/aggregates/user"
	"fmt"

	"github.com/google/uuid"
)

func NewCourse(teacher user.User, name string) (*Course, error) {
	if !teacher.IsTeacher() {
		return nil, fmt.Errorf("Only teachers can create courses")
	}

	nameVO, err := NewCourseName(name)
	if err != nil {
		return nil, err
	}

	return &Course{
		createdByID: teacher.ID(),
		name:        nameVO,
		isArchived:  false,
		id:          CourseID(uuid.New()),
		topicCount:  0,
		groupCount:  0,
	}, nil
}
