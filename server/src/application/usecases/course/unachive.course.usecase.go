package course_usercases

import (
	"adaptivetesting/src/application/dto/requests"
	"adaptivetesting/src/application/dto/responses"
	"adaptivetesting/src/application/mappers"
	"adaptivetesting/src/domain/aggregates/course"
	"adaptivetesting/src/domain/aggregates/user"
	"fmt"
)

type CourseUnarchive struct {
	courseRepo course.ICourseRepository
}

func (this CourseUnarchive) Execute(current_user user.User, data requests.ArchiveCourseDTO) (*responses.CourseResponseWithUser, error) {
	crs, err := this.courseRepo.GetById(data.CourseID)
	if err != nil {
		return nil, err
	}

	if !crs.IsUserCreator(current_user) {
		return nil, fmt.Errorf("You not a course creator")
	}

	if err := crs.Activate(); err != nil {
		return nil, err
	}

	if err := this.courseRepo.Save(crs); err != nil {
		return nil, err
	}

	return mappers.CourseToResponseWithUser(crs, &current_user), nil
}
