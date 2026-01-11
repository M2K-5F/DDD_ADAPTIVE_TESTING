package interfaces

import (
	"adaptivetesting/src/domain/aggregates/topic"
	"adaptivetesting/src/domain/identificators"
	"context"
)

type ITopicReader interface {
	GetByID(ctx context.Context, topicID identificators.TopicID) (*topic.Topic, error)
}
