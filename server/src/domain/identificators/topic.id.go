package identificators

import "github.com/google/uuid"

type TopicID uuid.UUID

func (id TopicID) String() string {
	return uuid.UUID(id).String()
}

func NewTopicID() TopicID {
	return TopicID(uuid.New())
}

func ParseTopicID(plain string) (TopicID, error) {
	id, err := uuid.Parse(plain)
	if err != nil {
		return TopicID(uuid.Nil), err
	}
	return TopicID(id), nil
}
