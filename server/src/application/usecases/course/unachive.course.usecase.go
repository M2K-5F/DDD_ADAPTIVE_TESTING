package course_usercases

import (
	"adaptivetesting/src/application/dto"
	"adaptivetesting/src/application/mappers"
	"adaptivetesting/src/domain/aggregates/course"
	"adaptivetesting/src/domain/aggregates/user"
	"fmt"
)

type UnarchiveCourseUC struct {
	courseRepo course.ICourseRepository
}

func (this UnarchiveCourseUC) Execute(current_user user.User, data dto.ArchiveCourseDTO) (*dto.CourseResponseWithUser, error) {
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
