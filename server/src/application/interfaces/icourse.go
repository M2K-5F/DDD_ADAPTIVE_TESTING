package interfaces

import (
	"adaptivetesting/src/domain/aggregates/course"
	"adaptivetesting/src/domain/identificators"
	"context"
)

type ICourseReader interface {
	GetByID(ctx context.Context, id identificators.CourseID) (*course.Course, error)
	GetByName(ctx context.Context, name string) (*course.Course, error)
	SelectCreatedByUser(ctx context.Context, userID identificators.UserID) ([]*course.Course, error)
}
