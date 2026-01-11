package enrollment

import (
	"adaptivetesting/src/domain/identificators"
)

type TopicProgress struct {
	maxScore float32
	attempts []*TopicAttempt
	topicID  identificators.TopicID
}

func (this *TopicProgress) registerAttempt(score float32) error {
	attempt, err := NewTopicAttempt(score)
	if err != nil {
		return err
	}

	if score > this.maxScore {
		this.maxScore = score
	}

	this.attempts = append(this.attempts, attempt)
	return nil
}

func NewTopicProgress(topicID identificators.TopicID) (*TopicProgress, error) {
	return &TopicProgress{
		maxScore: 0,
		topicID:  topicID,
		attempts: []*TopicAttempt{},
	}, nil
}
