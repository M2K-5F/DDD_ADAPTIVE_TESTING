package enrollment

import (
	"adaptivetesting/src/domain/identificators"
)

type EnrollmentFabric struct{}

func (EnrollmentFabric) CreateEnrollment(
	groupID identificators.GroupID,
	courseID identificators.CourseID,
	userID identificators.UserID,
) (
	*Enrollment,
	error,
) {
	return &Enrollment{
		id:              identificators.NewEnrollmentID(),
		groupID:         groupID,
		courseID:        courseID,
		userID:          userID,
		topicProgresses: []*TopicProgress{},
		progress:        0,
	}, nil
}
