package enrollment

import (
	"adaptivetesting/src/domain/identificators"
	"fmt"
)

func (this *Enrollment) RegisterAttempt(topicID identificators.TopicID, score float32) error {
	topicProgress, err := this.TopicProgressByID(topicID)
	if err != nil {
		this.startTopic(topicID)
	}

	topicProgress, err = this.TopicProgressByID(topicID)
	if err != nil {
		return err
	}

	err = topicProgress.registerAttempt(score)
	if err != nil {
		return err
	}

	return nil
}

func (this *Enrollment) startTopic(topicID identificators.TopicID) error {
	if _, err := this.TopicProgressByID(topicID); err == nil {
		return fmt.Errorf("This topic (%d) already started", topicID)
	}

	progress, err := NewTopicProgress(topicID)
	if err != nil {
		return err
	}

	this.topicProgresses = append(this.topicProgresses, progress)

	return nil
}
