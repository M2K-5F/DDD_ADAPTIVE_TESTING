package mappers

import (
	"adaptivetesting/src/application/dto/responses"
	"adaptivetesting/src/domain/aggregates/course"
	"adaptivetesting/src/domain/aggregates/topic"
	"adaptivetesting/src/domain/aggregates/user"
)

func TopicToResponse(topic *topic.Topic) *responses.TopicResponse {
	return &responses.TopicResponse{
		ID:          topic.ID().String(),
		Name:        topic.Name().String(),
		ByCourseID:  topic.ByCourseID().String(),
		CreatedByID: topic.CreatedById().String(),
	}
}

func TopicToNestedResponse(topic *topic.Topic, course *course.Course, creator *user.User) *responses.TopicNestedResponse {
	return &responses.TopicNestedResponse{
		ID:        topic.ID().String(),
		Name:      topic.Name().String(),
		ByCourse:  *CourseToResponse(course),
		CreatedBy: *UserToResponse(creator),
	}
}
