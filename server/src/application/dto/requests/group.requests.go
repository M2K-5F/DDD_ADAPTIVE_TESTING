package requests

import "adaptivetesting/src/domain/identificators"

type CreateGroupRequest struct {
	Name            string
	CourseID        identificators.CourseID
	MaxStudentCount int
}
