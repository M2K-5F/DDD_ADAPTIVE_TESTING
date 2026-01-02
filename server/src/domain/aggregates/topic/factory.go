package topic

import (
	"adaptivetesting/src/domain/aggregates/course"

	"github.com/google/uuid"
)

func NewTopic(name string, course course.Course) (*Topic, error) {
	nameVo, err := NewTopicName(name)

	if err != nil {
		return nil, err
	}

	return &Topic{
		id:            TopicID(uuid.New()),
		courseID:      course.ID(),
		name:          nameVo,
		isArchived:    false,
		questionCount: 0,
	}, nil
}
