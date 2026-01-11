package enrollment

import (
	"adaptivetesting/src/domain/identificators"
	"fmt"
)

type Enrollment struct {
	id              identificators.EnrollmentID
	groupID         identificators.GroupID
	courseID        identificators.CourseID
	userID          identificators.UserID
	topicProgresses []*TopicProgress
	progress        float32
}

func (e Enrollment) ID() identificators.EnrollmentID {
	return e.id
}

func (this *Enrollment) GroupID() identificators.GroupID {
	return this.groupID
}

func (this *Enrollment) CourseID() identificators.CourseID {
	return this.courseID
}

func (this *Enrollment) UserID() identificators.UserID {
	return this.userID
}

func (this *Enrollment) Progress() float32 {
	return this.progress
}

func (this *Enrollment) TopicProgressByID(topicID identificators.TopicID) (*TopicProgress, error) {
	for _, tp := range this.topicProgresses {
		if tp.topicID == topicID {
			return tp, nil
		}
	}

	return nil, fmt.Errorf("You dont pass topic with id (%d)", topicID)
}

func (this *Enrollment) TopicProgresses() []*TopicProgress {
	return this.topicProgresses
}
