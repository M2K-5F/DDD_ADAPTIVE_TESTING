package course_usercases

import (
	"adaptivetesting/src/application/dto/mappers"
	"adaptivetesting/src/application/dto/requests"
	"adaptivetesting/src/application/dto/responses"
	"adaptivetesting/src/application/interfaces"
	"context"
	"fmt"
)

type GetCoursesCreatedByUser struct {
	courseReader interfaces.ICourseReader
	userReader   interfaces.IUserReader
}

func (this *GetCoursesCreatedByUser) Execute(
	ctx context.Context,
	data *requests.GetCoursesByUserDTO,
) (
	*[]responses.CourseNestedResponse,
	error,
) {
	current_techer, err := this.userReader.GetByID(ctx, data.UserID)
	if err != nil {
		return nil, err
	}

	if !current_techer.IsTeacher() {
		return nil, fmt.Errorf("This user is not a teacher")
	}

	courses, err := this.courseReader.SelectCreatedByUser(ctx, data.UserID)
	if err != nil {
		return nil, err
	}

	var response []responses.CourseNestedResponse

	for _, crs := range courses {
		response = append(response, *mappers.CourseToNestedResponse(crs, current_techer))
	}

	return &response, nil
}
