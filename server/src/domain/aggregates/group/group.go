package group

import (
	"adaptivetesting/src/domain/aggregates/identificators"
)

type Group struct {
	id              identificators.GroupID
	byCourseID      identificators.CourseID
	createdByID     identificators.UserID
	name            GroupName
	studentCount    int
	maxStudentCount int
}

func (this *Group) ID() identificators.GroupID {
	return this.id
}

func (this *Group) ByCourseID() identificators.CourseID {
	return this.byCourseID
}

func (this *Group) CreatedByID() identificators.UserID {
	return this.createdByID
}

func (this *Group) Name() GroupName {
	return this.name
}

func (this *Group) StudentCount() int {
	return this.studentCount
}

func (this *Group) MaxStudentCount() int {
	return this.maxStudentCount
}
