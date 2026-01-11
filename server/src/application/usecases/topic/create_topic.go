package course_usercases

import (
	"adaptivetesting/src/application/dto/mappers"
	"adaptivetesting/src/application/dto/requests"
	"adaptivetesting/src/application/dto/responses"
	"adaptivetesting/src/application/interfaces"
	"adaptivetesting/src/domain/aggregates/topic"
	"adaptivetesting/src/domain/aggregates/user"
	"context"
	"fmt"
)

type CreateTopic struct {
	writer       interfaces.IWriter
	courseReader interfaces.ICourseReader
}

func (this CreateTopic) Execute(
	ctx context.Context,
	current_user *user.User,
	data requests.CreateTopicRequest,
) (
	*responses.CourseNestedResponse,
	error,
) {
	if !current_user.IsTeacher() {
		return nil, fmt.Errorf("Only teachers can create topics")
	}

	current_course, err := this.courseReader.GetByID(ctx, data.CourseId)
	if err != nil {
		return nil, err
	}

	if !current_course.IsCreatedBy(current_user.ID()) {
		return nil, fmt.Errorf("Only creator of this course can add topic to it")
	}

	createdTopic, err := topic.Fabric.CreateTopic(
		data.Name,
		current_user.ID(),
		data.CourseId,
	)

	if err := this.writer.Execute(ctx, func(writer interfaces.ITransactionWriter) error {
		return writer.SaveTopic(createdTopic)
	}); err != nil {
		return nil, err
	}

	return mappers.CourseToNestedResponse(current_course, current_user), nil
}
