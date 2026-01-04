package enrollment

import (
	"adaptivetesting/src/domain/aggregates/identificators"
)

type Enrollment struct {
	id            identificators.EnrollmentID
	groupID       identificators.GroupID
	courseID      identificators.CourseID
	userID        identificators.UserID
	topicProgress map[identificators.TopicID]*TopicProgress
	progress      float32
}

func (e Enrollment) ID() identificators.EnrollmentID {
	return e.id
}
