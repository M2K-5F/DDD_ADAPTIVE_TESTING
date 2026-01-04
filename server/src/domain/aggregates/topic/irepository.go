package topic

import "adaptivetesting/src/domain/aggregates/identificators"

type ITopicRepository interface {
	Save(*Topic) error
	GetByID(id identificators.TopicID) (*Topic, error)
}
