package requests

import "adaptivetesting/src/domain/aggregates/identificators"

type CreateGroupRequest struct {
	Name            string
	CourseID        identificators.CourseID
	MaxStudentCount int
}
