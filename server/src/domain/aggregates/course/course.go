package course

import (
	"adaptivetesting/src/domain/aggregates/user"
	"fmt"

	"github.com/google/uuid"
)

type CourseID uuid.UUID

type Course struct {
	id          CourseID
	createdByID user.UserID
	name        CourseName
	isArchived  bool

	topicCount int
	groupCount int
}

func (c *Course) ID() CourseID {
	return c.id
}

func (this *Course) AddTopic() error {
	this.topicCount += 1
	return nil
}

func (this *Course) RemoveTopic() error {
	if this.topicCount > 1 {
		this.topicCount -= 1
		return nil
	}
	return fmt.Errorf("course topic count is 0")
}

func (this *Course) AddGroup() error {
	this.groupCount += 1
	return nil
}

func (this *Course) RemoveGroup() error {
	if this.groupCount > 1 {
		this.groupCount -= 1
		return nil
	}
	return fmt.Errorf("course group count is 0")
}

func (this *Course) Activate() error {
	if this.isArchived {
		this.isArchived = false
		return nil
	}
	return fmt.Errorf("course already active")
}

func (this *Course) Archivate() error {
	if !this.isArchived {
		this.isArchived = true
		return nil
	}
	return fmt.Errorf("course already archived")
}
