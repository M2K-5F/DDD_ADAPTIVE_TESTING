package course_usercases

import (
	"adaptivetesting/src/application/dto/requests"
	"adaptivetesting/src/application/dto/responses"
	"adaptivetesting/src/application/mappers"
	"adaptivetesting/src/domain/aggregates/course"
	"adaptivetesting/src/domain/aggregates/user"
)

type CourseCreate struct {
	courseRepo course.ICourseRepository
}

func (uc *CourseCreate) Execute(current_user user.User, data requests.CreateCourseDTO) (*responses.CourseResponseWithUser, error) {
	crs, err := course.NewCourse(current_user, data.Name)
	if err != nil {
		return nil, err
	}

	uc.courseRepo.Save(crs)
	return mappers.CourseToResponseWithUser(crs, &current_user), nil
}
