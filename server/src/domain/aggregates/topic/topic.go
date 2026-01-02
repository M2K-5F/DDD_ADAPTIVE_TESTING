package topic

import (
	"adaptivetesting/src/domain/aggregates/course"
	"fmt"

	"github.com/google/uuid"
)

type TopicID uuid.UUID

type Topic struct {
	id         TopicID
	courseID   course.CourseID
	name       TopicName
	isArchived bool

	questionCount int
}

func (this *Topic) ID() TopicID {
	return this.id
}

func (this *Topic) Activate() error {
	if this.isArchived {
		this.isArchived = false
		return nil
	}
	return fmt.Errorf("topic already active")
}

func (this *Topic) Archivate() error {
	if !this.isArchived {
		this.isArchived = true
		return nil
	}
	return fmt.Errorf("topic already archived")
}
