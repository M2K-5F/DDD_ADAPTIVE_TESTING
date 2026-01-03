package course

import "adaptivetesting/src/domain/aggregates/user"

type ICourseRepository interface {
	Save(*Course) error
	GetById(id CourseID) (*Course, error)
	ListByUserID(userID user.UserID) ([]Course, error)
}
