package topic

import (
	"adaptivetesting/src/domain/identificators"
)

type TopicFabric struct{}

func (TopicFabric) CreateTopic(name string, userId identificators.UserID, courseID identificators.CourseID) (*Topic, error) {
	topicName, err := NewTopicName(name)
	if err != nil {
		return nil, err
	}

	return &Topic{
		byCourseID:  courseID,
		createdByID: userId,
		id:          identificators.NewTopicID(),
		name:        topicName,
		isArchived:  false,
	}, nil
}

func (TopicFabric) Recover(name, byCourseID, createdByID, id string, isArchived bool) (*Topic, error) {
	idVO, err := identificators.ParseTopicID(id)
	if err != nil {
		return nil, err
	}

	byCourseIDVO, err := identificators.ParseCourseID(byCourseID)
	if err != nil {
		return nil, err
	}

	createdByIDVO, err := identificators.ParseUserID(createdByID)
	if err != nil {
		return nil, err
	}

	nameVO, err := NewTopicName(name)
	if err != nil {
		return nil, err
	}

	return &Topic{
		id:          idVO,
		byCourseID:  byCourseIDVO,
		createdByID: createdByIDVO,
		name:        nameVO,
		isArchived:  isArchived,
	}, nil
}

var Fabric = TopicFabric{}
