package requests

import (
	"adaptivetesting/src/domain/aggregates/identificators"
)

type CreateTopicRequest struct {
	CourseId identificators.CourseID
	Name     string
}
