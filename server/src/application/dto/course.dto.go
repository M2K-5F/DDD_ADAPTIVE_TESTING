package dto

import (
	"adaptivetesting/src/domain/aggregates/course"
	"adaptivetesting/src/domain/aggregates/user"
)

type CreateCourseDTO struct {
	Name string
}

type ArchiveCourseDTO struct {
	CourseID course.CourseID
}

type GetCoursesByUserDTO struct {
	UserID user.UserID
}
