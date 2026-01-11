package course

import (
	"adaptivetesting/src/domain/identificators"
)

type CourseFabric struct{}

func (CourseFabric) CreateCourse(teacherID identificators.UserID, name string) (*Course, error) {

	nameVO, err := NewCourseName(name)
	if err != nil {
		return nil, err
	}

	return &Course{
		createdByID: teacherID,
		name:        nameVO,
		isArchived:  false,
		id:          identificators.NewCourseID(),
	}, nil
}

func (CourseFabric) Recover(id string, createdByID string, name string, isArchived bool) (*Course, error) {
	idVO, err := identificators.ParseCourseID(id)
	if err != nil {
		return nil, err
	}

	createdByIDVO, err := identificators.ParseUserID(createdByID)
	if err != nil {
		return nil, err
	}
	nameVO, err := NewCourseName(name)
	if err != nil {
		return nil, err
	}

	return &Course{
		id:          idVO,
		createdByID: createdByIDVO,
		isArchived:  isArchived,
		name:        nameVO,
	}, nil
}

var Fabric = CourseFabric{}
