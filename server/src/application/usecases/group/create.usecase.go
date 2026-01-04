package group_usecases

import (
	"adaptivetesting/src/application/dto/requests"
	"adaptivetesting/src/application/dto/responses"
	"adaptivetesting/src/application/mappers"
	"adaptivetesting/src/domain/aggregates/course"
	"adaptivetesting/src/domain/aggregates/group"
	"adaptivetesting/src/domain/aggregates/user"
	"fmt"
)

type CreateGroup struct {
	groupRepo  group.IRepository
	courseRepo course.ICourseRepository
}

func (this *CreateGroup) Execute(current_user *user.User, data *requests.CreateGroupRequest) (*responses.GroupNestedResponse, error) {
	current_course, err := this.courseRepo.GetById(data.CourseID)
	if err != nil {
		return nil, err
	}
	if !current_course.IsUserCreator(current_user) {
		return nil, fmt.Errorf("Only course creator can create a group")
	}

	created_group, err := group.NewGroup(*current_course, data.Name, data.MaxStudentCount)
	if err != nil {
		return nil, err
	}

	if err := current_course.AddGroup(created_group.ID()); err != nil {
		return nil, err
	}

	if err := this.groupRepo.Save(created_group); err != nil {
		return nil, err
	}

	if err := this.courseRepo.Save(current_course); err != nil {
		return nil, err
	}

	return mappers.GroupToNestedResponse(created_group, current_user, current_course), nil
}
