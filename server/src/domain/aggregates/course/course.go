package course

import (
	"adaptivetesting/src/domain/aggregates/user"
	"fmt"

	"github.com/google/uuid"
)

type CourseID uuid.UUID

func (id CourseID) String() string {
	return uuid.UUID(id).String()
}

type Course struct {
	id          CourseID
	createdByID user.UserID
	name        CourseName
	isArchived  bool

	topicCount int
	groupCount int
}

// GET
func (c *Course) ID() CourseID {
	return c.id
}

func (c *Course) Name() CourseName {
	return c.name
}

func (c *Course) CreatedByID() user.UserID {
	return c.createdByID
}

func (c *Course) IsArchived() bool {
	return c.isArchived
}

func (c *Course) TopicCount() int {
	return c.topicCount
}

func (c *Course) GroupCount() int {
	return c.groupCount
}

// SET
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

func (this *Course) IsUserCreator(user user.User) bool {
	return this.createdByID == user.ID()
}
