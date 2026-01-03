package course_usercases

import (
	"adaptivetesting/src/application/dto"
	"adaptivetesting/src/application/mappers"
	"adaptivetesting/src/domain/aggregates/course"
	"adaptivetesting/src/domain/aggregates/user"
	"fmt"
)

type ArchiveCourseUC struct {
	courseRepo course.ICourseRepository
}

func (uc *ArchiveCourseUC) Execute(current_user user.User, data *dto.ArchiveCourseDTO) (*dto.CourseResponseWithUser, error) {
	crs, err := uc.courseRepo.GetById(data.CourseID)
	if err != nil {
		return nil, err
	}

	if !crs.IsUserCreator(current_user) {
		return nil, fmt.Errorf("You not a creator of this course")
	}

	if err := crs.Archivate(); err != nil {
		return nil, err
	}

	err = uc.courseRepo.Save(crs)
	if err != nil {
		return nil, err
	}

	return mappers.CourseToResponseWithUser(crs, &current_user), nil
}
