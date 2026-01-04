package mappers

import (
	"adaptivetesting/src/application/dto/responses"
	"adaptivetesting/src/domain/aggregates/course"
	"adaptivetesting/src/domain/aggregates/group"
	"adaptivetesting/src/domain/aggregates/user"
)

func GroupToResponse(group *group.Group) *responses.GroupResponse {
	return &responses.GroupResponse{
		ID:           group.ID().String(),
		Name:         group.Name().String(),
		CreatedByID:  group.CreatedByID().String(),
		ByCourseID:   group.ByCourseID().String(),
		StudentCount: group.StudentCount(),
	}
}

func GroupToNestedResponse(group *group.Group, creator *user.User, course *course.Course) *responses.GroupNestedResponse {
	return &responses.GroupNestedResponse{
		ID:           group.ID().String(),
		Name:         group.Name().String(),
		CreatedBy:    *UserToDTO(creator),
		ByCourse:     *CourseToResponse(course),
		StudentCount: group.StudentCount(),
	}
}
