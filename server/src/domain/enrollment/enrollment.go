package enrollment

import (
	"adaptivetesting/src/domain/course"
	"adaptivetesting/src/domain/group"
	"adaptivetesting/src/domain/topic"
	"adaptivetesting/src/domain/user"
	"time"

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

type TopicProgress struct {
	maxScore float32
	attempts []TopicAttempt
}

type TopicAttempt struct {
	score       float32
	attemptedAt time.Time
}
