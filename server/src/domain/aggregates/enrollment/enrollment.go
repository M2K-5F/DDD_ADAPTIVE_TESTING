package enrollment

import (
	"adaptivetesting/src/domain/aggregates/course"
	"adaptivetesting/src/domain/aggregates/group"
	"adaptivetesting/src/domain/aggregates/topic"
	"adaptivetesting/src/domain/aggregates/user"

	"github.com/google/uuid"
)

type EnrollmentID uuid.UUID

type Enrollment struct {
	id            EnrollmentID
	groupID       group.GroupID
	courseID      course.CourseID
	userID        user.UserID
	topicProgress map[topic.TopicID]*TopicProgress
	progress      float32
}

func (e Enrollment) ID() EnrollmentID {
	return e.id
}
