package course_usercases

import (
	"adaptivetesting/src/application/dto/mappers"
	"adaptivetesting/src/application/dto/requests"
	"adaptivetesting/src/application/dto/responses"
	"adaptivetesting/src/application/interfaces"
	"adaptivetesting/src/domain/aggregates/user"
	"context"
	"fmt"
)

type CourseAchive struct {
	courseReader interfaces.ICourseReader
	writer       interfaces.IWriter
}

func (this *CourseAchive) Execute(
	ctx context.Context,
	current_user *user.User,
	data *requests.ArchiveCourseDTO,
) (
	*responses.CourseNestedResponse,
	error,
) {
	crs, err := this.courseReader.GetByID(ctx, data.CourseID)
	if err != nil {
		return nil, err
	}

	if !crs.IsCreatedBy(current_user.ID()) {
		return nil, fmt.Errorf("You not a creator of this course")
	}

	if err := crs.Archivate(); err != nil {
		return nil, err
	}

	if err := this.writer.Execute(ctx, func(writer interfaces.ITransactionWriter) error {
		return writer.SaveCourse(crs)
	}); err != nil {
		return nil, err
	}

	return mappers.CourseToNestedResponse(crs, current_user), nil
}
