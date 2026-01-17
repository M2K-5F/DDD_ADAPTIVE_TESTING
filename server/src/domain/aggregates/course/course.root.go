package course

import (
	"adaptivetesting/src/domain/identificators"
)

type Course struct {
	id          identificators.CourseID
	createdByID identificators.UserID
	name        CourseName
	isArchived  bool
}

func (c *Course) ID() identificators.CourseID {
	return c.id
}

func (c *Course) Name() CourseName {
	return c.name
}

func (c *Course) CreatedByID() identificators.UserID {
	return c.createdByID
}

func (c *Course) IsArchived() bool {
	return c.isArchived
}

type CoursePersistense struct {
	ID          string
	CreatedByID string
	Name        string
	IsArchived  bool
}

func (this *Course) ToPersistense() *CoursePersistense {
	return &CoursePersistense{
		ID:          this.id.String(),
		Name:        this.name.String(),
		CreatedByID: this.createdByID.String(),
		IsArchived:  this.isArchived,
	}
}
