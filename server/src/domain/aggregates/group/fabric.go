package group

import (
	"adaptivetesting/src/domain/identificators"
	"fmt"
)

type GroupFabric struct{}

func (GroupFabric) CreateGroup(createdByID identificators.UserID, courseID identificators.CourseID, name string, maxStudentCount int) (*Group, error) {
	nameVO, err := NewGroupName(name)
	if err != nil {
		return nil, err
	}

	return &Group{
		id:              identificators.NewGroupID(),
		name:            nameVO,
		byCourseID:      courseID,
		createdByID:     createdByID,
		maxStudentCount: maxStudentCount,
		studentCount:    0,
	}, nil
}

func (GroupFabric) Recover(
	createdByID, byCourseID, id, name string, maxStudentCount, studentCount int,
) (*Group, error) {
	idVO, err := identificators.ParseGroupID(id)
	if err != nil {
		return nil, err
	}

	nameVO, err := NewGroupName(name)
	if err != nil {
		return nil, err
	}

	createdByIDVO, err := identificators.ParseUserID(createdByID)
	if err != nil {
		return nil, err
	}

	byCourseIDVO, err := identificators.ParseCourseID(byCourseID)
	if err != nil {
		return nil, err
	}

	if maxStudentCount < 0 || studentCount < 0 || studentCount > maxStudentCount {
		return nil, fmt.Errorf("Unvalid student count")
	}

	return &Group{
		id:              idVO,
		createdByID:     createdByIDVO,
		byCourseID:      byCourseIDVO,
		name:            nameVO,
		maxStudentCount: maxStudentCount,
		studentCount:    studentCount,
	}, nil
}
