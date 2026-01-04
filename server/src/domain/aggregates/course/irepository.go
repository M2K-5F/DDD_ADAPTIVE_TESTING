package course

import (
	"adaptivetesting/src/domain/aggregates/identificators"
)

type ICourseRepository interface {
	Save(*Course) error
	GetById(id identificators.CourseID) (*Course, error)
	ListByUserID(userID identificators.UserID) ([]Course, error)
}
