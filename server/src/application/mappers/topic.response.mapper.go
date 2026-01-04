package mappers

import (
	"adaptivetesting/src/application/dto/responses"
	"adaptivetesting/src/domain/aggregates/course"
	"adaptivetesting/src/domain/aggregates/topic"
	"adaptivetesting/src/domain/aggregates/user"
)

func TopicToNestedResponse(topic *topic.Topic, course *course.Course, user *user.User) *responses.TopicResponseWithCourseAndUser {
	return &responses.TopicResponseWithCourseAndUser{
		ID:        topic.ID().String(),
		Name:      topic.Name().String(),
		ByCourse:  *CourseToResponse(course),
		CreatedBy: *UserToDTO(user),
	}
}

func TopicToResponse(topic *topic.Topic) *responses.TopicResponse {
	return &responses.TopicResponse{
		ID:          topic.ID().String(),
		Name:        topic.Name().String(),
		ByCourseID:  topic.ByCourseID().String(),
		CreatedByID: topic.CreatedById().String(),
	}
}
