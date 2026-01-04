package topic

import (
	"adaptivetesting/src/domain/aggregates/identificators"
	"adaptivetesting/src/domain/aggregates/user"
	"fmt"
)

type Topic struct {
	id          identificators.TopicID
	courseID    identificators.CourseID
	createdByID identificators.UserID
	name        TopicName
	isArchived  bool

	questionCount int
	questionIDs   []identificators.QuestionID
}

// QUERY
func (this *Topic) ID() identificators.TopicID {
	return this.id
}

func (this *Topic) Name() TopicName {
	return this.name
}

func (this *Topic) ByCourseID() identificators.CourseID {
	return this.courseID
}

func (this *Topic) CreatedById() identificators.UserID {
	return this.createdByID
}

func (this *Topic) IsCreatedBy(user *user.User) bool {
	return this.createdByID == user.ID()
}

// COMMAND
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

func (this *Topic) AddQuestion() error {
	this.questionCount += 1
	return nil
}

func (this *Topic) RemoveQuestion() error {
	if this.questionCount == 0 {
		return fmt.Errorf("This topic are not have questions")
	}

	this.questionCount -= 0
	return nil
}
