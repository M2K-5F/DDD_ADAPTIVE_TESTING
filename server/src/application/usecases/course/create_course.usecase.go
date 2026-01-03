package course_usercases

import (
	"adaptivetesting/src/application/dto"
	"adaptivetesting/src/application/mappers"
	"adaptivetesting/src/domain/aggregates/course"
	"adaptivetesting/src/domain/aggregates/user"
)

type CreateCourseUC struct {
	courseRepo course.ICourseRepository
}

func (uc *CreateCourseUC) Execute(current_user user.User, data dto.CreateCourseDTO) (*dto.CourseResponseWithUser, error) {
	crs, err := course.NewCourse(current_user, data.Name)
	if err != nil {
		return nil, err
	}

	uc.courseRepo.Save(crs)
	return mappers.CourseToResponseWithUser(crs, &current_user), nil
}
