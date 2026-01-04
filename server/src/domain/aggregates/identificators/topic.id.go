package identificators

import "github.com/google/uuid"

type TopicID uuid.UUID

func (id TopicID) String() string {
	return uuid.UUID(id).String()
}
