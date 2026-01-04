package course_usercases

import (
	"adaptivetesting/src/application/dto/requests"
	"adaptivetesting/src/application/dto/responses"
	"adaptivetesting/src/application/mappers"
	"adaptivetesting/src/domain/aggregates/course"
	"adaptivetesting/src/domain/aggregates/user"
	"fmt"
)

type GetCoursesCreatedByUser struct {
	courseRepo course.ICourseRepository
	userRepo   user.IUserRepository
}

func (this *GetCoursesCreatedByUser) Execute(data *requests.GetCoursesByUserDTO) (*[]responses.CourseResponseWithUser, error) {

	current_techer, err := this.userRepo.GetByID(data.UserID)
	if err != nil {
		return nil, err
	}

	if !current_techer.IsTeacher() {
		return nil, fmt.Errorf("This user is not a teacher")
	}

	courses, err := this.courseRepo.ListByUserID(data.UserID)
	if err != nil {
		return nil, err
	}

	var response []responses.CourseResponseWithUser

	for _, crs := range courses {
		response = append(response, *mappers.CourseToResponseWithUser(&crs, current_techer))
	}

	return &response, nil
}
