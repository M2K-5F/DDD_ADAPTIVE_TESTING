package course_usercases

import (
	"adaptivetesting/src/application/dto/mappers"
	"adaptivetesting/src/application/dto/requests"
	"adaptivetesting/src/application/dto/responses"
	"context"
	"fmt"
)

func (this *CourseUseCases) GetByCreator(
	ctx context.Context,
	data *requests.GetCoursesByUserDTO,
) (
	*[]responses.CourseNestedResponse,
	error,
) {
	current_techer, err := this.deps.UserRdr.GetByID(ctx, data.UserID)
	if err != nil {
		return nil, err
	}

	if !current_techer.IsTeacher() {
		return nil, fmt.Errorf("This user is not a teacher")
	}

	courses, err := this.deps.CourseRdr.SelectCreatedByUser(ctx, data.UserID)
	if err != nil {
		return nil, err
	}

	var response []responses.CourseNestedResponse

	for _, crs := range courses {
		response = append(response, *mappers.CourseToNestedResponse(crs, current_techer))
	}

	return &response, nil
}
