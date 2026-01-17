package course_usercases

import (
	"adaptivetesting/src/application/dto/mappers"
	"adaptivetesting/src/application/dto/requests"
	"adaptivetesting/src/application/dto/responses"
	"adaptivetesting/src/application/interfaces"
	"adaptivetesting/src/domain/aggregates/course"
	"adaptivetesting/src/domain/aggregates/user"
	"context"
)

func (this *CourseUseCases) Create(
	ctx context.Context,
	current_user user.User,
	data requests.CreateCourseDTO,
) (
	*responses.CourseNestedResponse,
	error,
) {
	crs, err := course.Fabric.CreateCourse(current_user.ID(), data.Name)
	if err != nil {
		return nil, err
	}

	if err := this.deps.Writer.Execute(ctx, func(writer interfaces.ITransactionWriter) error {
		return writer.SaveCourse(crs)
	}); err != nil {
		return nil, err
	}

	return mappers.CourseToNestedResponse(crs, &current_user), nil
}
