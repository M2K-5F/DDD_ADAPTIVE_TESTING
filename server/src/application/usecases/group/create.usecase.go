package group_usecases

import (
	"adaptivetesting/src/application/dto/mappers"
	"adaptivetesting/src/application/dto/requests"
	"adaptivetesting/src/application/dto/responses"
	"adaptivetesting/src/application/interfaces"
	"adaptivetesting/src/domain/aggregates/group"
	"adaptivetesting/src/domain/aggregates/user"
	"context"
	"fmt"
)

type CreateGroup struct {
	courseReader interfaces.ICourseReader
	writer       interfaces.IWriter
}

func (this *CreateGroup) Execute(
	ctx context.Context,
	current_user *user.User,
	data *requests.CreateGroupRequest,
) (
	*responses.GroupNestedResponse,
	error,
) {
	current_course, err := this.courseReader.GetByID(ctx, data.CourseID)
	if err != nil {
		return nil, err
	}
	if !current_course.IsCreatedBy(current_user.ID()) {
		return nil, fmt.Errorf("Only course creator can create a group")
	}

	created_group, err := group.Fabric.CreateGroup(current_user.ID(), current_course.ID(), data.Name, data.MaxStudentCount)
	if err != nil {
		return nil, err
	}

	if err := this.writer.Execute(ctx, func(writer interfaces.ITransactionWriter) error {
		return writer.SaveGroup(created_group)
	}); err != nil {
		return nil, err
	}

	return mappers.GroupToNestedResponse(created_group, current_user, current_course), nil
}
