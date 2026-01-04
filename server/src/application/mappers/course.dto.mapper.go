package mappers

import (
	"adaptivetesting/src/application/dto/responses"
	"adaptivetesting/src/domain/aggregates/course"
	"adaptivetesting/src/domain/aggregates/user"
)

func CourseToResponse(course *course.Course) *responses.CourseResponse {
	return &responses.CourseResponse{
		ID:          course.ID().String(),
		Name:        course.Name().String(),
		CreatedByID: course.CreatedByID().String(),
	}
}

func CourseToResponseWithUser(course *course.Course, creator *user.User) *responses.CourseResponseWithUser {
	return &responses.CourseResponseWithUser{
		ID:        course.ID().String(),
		Name:      course.Name().String(),
		CreatedBy: *UserToDTO(creator),
	}
}
