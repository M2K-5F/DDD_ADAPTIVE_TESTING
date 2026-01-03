package mappers

import (
	"adaptivetesting/src/application/dto"
	"adaptivetesting/src/domain/aggregates/course"
	"adaptivetesting/src/domain/aggregates/user"
)

func CourseToResponse(course *course.Course) *dto.CourseResponse {
	return &dto.CourseResponse{
		ID:          course.ID().String(),
		Name:        course.Name().String(),
		CreatedByID: course.CreatedByID().String(),
	}
}

func CourseToResponseWithUser(course *course.Course, creator *user.User) *dto.CourseResponseWithUser {
	return &dto.CourseResponseWithUser{
		ID:        course.ID().String(),
		Name:      course.Name().String(),
		CreatedBy: *UserToDTO(creator),
	}
}
