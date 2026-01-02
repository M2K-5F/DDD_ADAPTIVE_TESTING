package enrollment

import (
	"fmt"
	"time"
)

type TopicAttempt struct {
	score       float32
	attemptedAt time.Time
}

func NewTopicAttempt(score float32) (*TopicAttempt, error) {
	if score < 0 || score > 1 {
		return nil, fmt.Errorf("Invalid score value: %g", score)
	}

	return &TopicAttempt{
		score:       score,
		attemptedAt: time.Now(),
	}, nil
}
