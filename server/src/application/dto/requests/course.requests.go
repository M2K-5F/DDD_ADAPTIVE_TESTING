package requests

import (
	"adaptivetesting/src/domain/identificators"
)

type CreateCourseDTO struct {
	Name string
}

type ArchiveCourseDTO struct {
	CourseID identificators.CourseID
}

type GetCoursesByUserDTO struct {
	UserID identificators.UserID
}
