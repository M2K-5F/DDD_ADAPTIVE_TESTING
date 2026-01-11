package interfaces

import (
	"adaptivetesting/src/domain/aggregates/group"
	"adaptivetesting/src/domain/identificators"
	"context"
)

type IGroupReader interface {
	GetByID(ctx context.Context, id identificators.GroupID) (*group.Group, error)
	ListByCourse(ctx context.Context, courseID identificators.CourseID) ([]*group.Group, error)
}
