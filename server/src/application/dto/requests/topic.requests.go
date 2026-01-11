package requests

import (
	"adaptivetesting/src/domain/identificators"
)

type CreateTopicRequest struct {
	CourseId identificators.CourseID
	Name     string
}
