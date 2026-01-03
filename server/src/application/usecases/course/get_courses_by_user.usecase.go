package course_usercases

import (
	"adaptivetesting/src/application/dto"
	"adaptivetesting/src/application/mappers"
	"adaptivetesting/src/domain/aggregates/course"
	"adaptivetesting/src/domain/aggregates/user"
	"fmt"
)

type GetCoursesByUserUC struct {
	courseRepo course.ICourseRepository
	userRepo   user.IUserRepository
}

func (this *GetCoursesByUserUC) Execute(data *dto.GetCoursesByUserDTO) (*[]dto.CourseResponseWithUser, error) {

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

	var response []dto.CourseResponseWithUser

	for _, crs := range courses {
		response = append(response, *mappers.CourseToResponseWithUser(&crs, current_techer))
	}

	return &response, nil
}
