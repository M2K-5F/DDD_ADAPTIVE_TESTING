package enrollment

import (
	"adaptivetesting/src/domain/aggregates/course"
	"adaptivetesting/src/domain/aggregates/group"
	"adaptivetesting/src/domain/aggregates/topic"
	"adaptivetesting/src/domain/aggregates/user"

	"github.com/google/uuid"
)

func NewEnrollment(group group.Group, course course.Course, user user.User) *Enrollment {
	return &Enrollment{
		id:            EnrollmentID(uuid.New()),
		groupID:       group.ID(),
		courseID:      course.ID(),
		userID:        user.ID(),
		topicProgress: map[topic.TopicID]*TopicProgress{},
		progress:      0,
	}
}
