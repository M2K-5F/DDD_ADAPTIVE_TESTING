package enrollment

type TopicProgress struct {
	maxScore float32
	attempts []TopicAttempt
}

func NewTopicProgress() (*TopicProgress, error) {
	return &TopicProgress{
		maxScore: 0,
		attempts: []TopicAttempt{},
	}, nil
}
