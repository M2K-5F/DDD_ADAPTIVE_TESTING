package topic

import (
	"adaptivetesting/src/domain/aggregates/course"
	"adaptivetesting/src/domain/aggregates/identificators"
	"adaptivetesting/src/domain/aggregates/user"

	"github.com/google/uuid"
)

func NewTopic(name string, course course.Course, teacher user.User) (*Topic, error) {
	nameVo, err := NewTopicName(name)

	if err != nil {
		return nil, err
	}

	return &Topic{
		id:            identificators.TopicID(uuid.New()),
		courseID:      course.ID(),
		createdByID:   teacher.ID(),
		name:          nameVo,
		isArchived:    false,
		questionCount: 0,
	}, nil
}
