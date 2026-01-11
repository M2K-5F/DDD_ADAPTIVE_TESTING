package topic

import (
	"adaptivetesting/src/domain/identificators"
)

type Topic struct {
	id          identificators.TopicID
	byCourseID  identificators.CourseID
	createdByID identificators.UserID
	name        TopicName
	isArchived  bool
}

func (this *Topic) ID() identificators.TopicID {
	return this.id
}

func (this *Topic) Name() TopicName {
	return this.name
}

func (this *Topic) ByCourseID() identificators.CourseID {
	return this.byCourseID
}

func (this *Topic) CreatedById() identificators.UserID {
	return this.createdByID
}
